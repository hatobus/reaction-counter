.PHONY: upload
upload:
	gcloud functions deploy $(FUNCITON_NAME) \
		--runtime go113 \
		--trigger-http \
		-set-env-vers "ACCESS_TOKEN=$(ACCESS_TOKEN),VERIFICATION_TOKEN=$(VERIFICATION_TOKEN)" \
		--allow-unauthenticated

