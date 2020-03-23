SHELL		:= /bin/bash
NAMESPACE	:= default

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
	kubectl apply -f https://raw.githubusercontent.com/intel/multus-cni/master/images/multus-daemonset.yml

install: setup ## Install all resources (CR/CRD's, RBAC and Operator)
	$(call echo_green," ....... Creating namespace .......")
	-kubectl create namespace ${NAMESPACE}
	$(call echo_green," ....... Applying CRDs .......")
	kubectl apply -f deploy/crds/bans.io_bandwidthslice_crd.yaml -n ${NAMESPACE}
	$(call echo_green," ....... Applying Rules and Service Account .......")
	kubectl apply -f deploy/role.yaml -n ${NAMESPACE}
	kubectl apply -f deploy/role_binding.yaml -n ${NAMESPACE}
	kubectl apply -f deploy/service_account.yaml -n ${NAMESPACE}
	$(call echo_green," ....... Applying Operator .......")
	kubectl apply -f deploy/operator.yaml -n ${NAMESPACE}
	${SHELL} scripts/wait_pods_running.sh ${NAMESPACE}
	# $(call echo_green," ....... Creating the CRs .......")
	# kubectl apply -f deploy/crds/bans.io_v1alpha1_bandwidthslice_cr.yaml -n ${NAMESPACE}

uninstall: ## Uninstall all that all performed in the $ make install
	$(call echo_red," ....... Uninstalling .......")
	$(call echo_red," ....... Deleting CRDs.......")
	-kubectl delete -f deploy/crds/bans.io_bandwidthslice_crd.yaml -n ${NAMESPACE}
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
	operator-sdk build steven30801/onos-bandwidth-operator:bmv2-fabric
	$(call echo_green," ...... Pushing image ......")
	docker push steven30801/onos-bandwidth-operator:bmv2-fabric

reset-onos:
	-helm uninstall onos
	-helm uninstall mininet
	-helm uninstall activate-bw-mgnt
	${SHELL} scripts/wait_pods_terminating.sh ${NAMESPACE}
