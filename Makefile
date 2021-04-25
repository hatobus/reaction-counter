.PHONY: upload
upload:
	gcloud functions deploy $(FUNCTION_NAME) \
		--runtime go113 \
		--trigger-http \
		--set-env-vars "ACCESS_TOKEN=$(ACCESS_TOKEN),VERIFICATION_TOKEN=$(VERIFICATION_TOKEN)" \
		--entry-point ReactionCounter \
		--allow-unauthenticated

