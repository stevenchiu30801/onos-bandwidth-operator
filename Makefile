SHELL		:= /bin/bash
NAMESPACE	:= default
MAKEDIR		:= $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
DEPLOY		:= $(MAKEDIR)/deploy
HELMDIR		:= $(MAKEDIR)/helm-charts

HELM_ARGS	?= --install --wait --timeout 6m -f $(HELMDIR)/configs/values.yaml

COLOR_WHITE			= \033[0m
COLOR_LIGHT_GREEN	= \033[1;32m
COLOR_LIGHT_RED		= \033[1;31m

define echo_green
	@echo -e "${COLOR_LIGHT_GREEN}${1}${COLOR_WHITE}"
endef

define echo_red
	@echo -e "${COLOR_LIGHT_RED}${1}${COLOR_WHITE}"
endef

.PHONY: setup install uninstall build reset-onos

setup: ## Setup environment
	$(call echo_green," ...... Setup Environment ......")
	# kubectl apply -f https://raw.githubusercontent.com/intel/multus-cni/master/images/multus-daemonset.yml

transport: ## Prepare transport network environment
	helm upgrade $(HELM_ARGS) onos $(HELMDIR)/onos
	@until http -a onos:rocks --ignore-stdin --check-status GET http://127.0.0.1:30181/onos/v1/applications/org.onosproject.drivers.barefoot-pro 2>&- | jq '.state' 2>&- | grep 'ACTIVE' >/dev/null; \
	do \
		echo "Waiting for ONOS to be ready"; \
		sleep 5; \
	done
	curl -u onos:rocks -X POST -H "Content-Type:application/json" -d @$(DEPLOY)/onos-device-netcfg.json http://127.0.0.1:30181/onos/v1/network/configuration
	@until http -a onos:rocks --ignore-stdin --check-status GET http://127.0.0.1:30181/onos/v1/devices/device:fabric:s1 2>&- | jq '.available' 2>&- | grep 'true' >/dev/null; \
	do \
		echo "Waiting for P4 device to be connected"; \
		sleep 3; \
	done
	curl -u onos:rocks -X POST -H "Content-Type:application/json" -d @$(DEPLOY)/onos-queue-netcfg.json http://127.0.0.1:30181/onos/v1/network/configuration
	helm upgrade $(HELM_ARGS) activate-bw-mgnt $(HELMDIR)/onos-app
	@until kubectl get job -o=jsonpath='{.items[?(@.status.succeeded==1)].metadata.name}' | grep 'activate-bw-mgnt-onos-app' >/dev/null; \
	do \
		echo "Waiting for bandwidth management application to be activated"; \
		sleep 3; \
	done
	@until ! http -a onos:rocks GET http://127.0.0.1:30181/onos/v1/flows/device:fabric:s1 2>&- | jq '.flows[].state' | grep 'PENDING_ADD' >/dev/null; \
	do \
		echo "Waiting for flows of bandwidth management to be added"; \
		sleep 5; \
	done

install: setup transport ## Install all resources (CR/CRD's, RBAC and Operator)
	$(call echo_green," ....... Creating namespace .......")
	-kubectl create namespace ${NAMESPACE}
	$(call echo_green," ....... Applying CRDs .......")
	kubectl apply -f deploy/crds/bans.io_bandwidthslice_crd.yaml -n ${NAMESPACE}
	kubectl apply -f deploy/crds/bans.io_fabricconfigs_crd.yaml -n ${NAMESPACE}
	$(call echo_green," ....... Applying Rules and Service Account .......")
	kubectl apply -f deploy/role.yaml -n ${NAMESPACE}
	kubectl apply -f deploy/role_binding.yaml -n ${NAMESPACE}
	kubectl apply -f deploy/service_account.yaml -n ${NAMESPACE}
	$(call echo_green," ....... Applying Operator .......")
	kubectl apply -f deploy/operator.yaml -n ${NAMESPACE}
	# ${SHELL} scripts/wait_pods_running.sh ${NAMESPACE}
	# $(call echo_green," ....... Creating the CRs .......")
	# kubectl apply -f deploy/crds/bans.io_v1alpha1_bandwidthslice_cr.yaml -n ${NAMESPACE}

uninstall: ## Uninstall all that all performed in the $ make install
	$(call echo_red," ....... Uninstalling .......")
	$(call echo_red," ....... Deleting CRDs.......")
	-kubectl delete -f deploy/crds/bans.io_bandwidthslice_crd.yaml -n ${NAMESPACE}
	# -kubectl delete -f deploy/crds/bans.io_fabricconfigs_crd.yaml -n ${NAMESPACE}
	$(call echo_red," ....... Deleting Rules and Service Account .......")
	-kubectl delete -f deploy/role.yaml -n ${NAMESPACE}
	-kubectl delete -f deploy/role_binding.yaml -n ${NAMESPACE}
	-kubectl delete -f deploy/service_account.yaml -n ${NAMESPACE}
	$(call echo_red," ....... Deleting Operator .......")
	-kubectl delete -f deploy/operator.yaml -n ${NAMESPACE}
	$(call echo_red," ....... Deleting namespace ${NAMESPACE}.......")
	-kubectl delete namespace ${NAMESPACE}

build: ## Build Operator
	$(call echo_green," ...... Building Operator ......")
	operator-sdk build steven30801/onos-bandwidth-operator:latest
	$(call echo_green," ...... Pushing image ......")
	docker push steven30801/onos-bandwidth-operator:latest

reset-onos:
	-helm uninstall onos
	-helm uninstall activate-bw-mgnt
	${SHELL} scripts/wait_pods_terminating.sh ${NAMESPACE}
