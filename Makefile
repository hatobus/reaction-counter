.PHONY: upload
upload:
	gcloud functions deploy $(FUNCITON_NAME) \
		--runtime go113 \
		--trigger-http \
		-set-env-vers "SLACK_SECRET=$(SLACK_SIGNING_SECRET),KG_API_KEY=$(KG_API_KEY)" \
		--allow-unauthenticated

