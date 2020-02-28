build: 
		docker build . -t tarunpothulapati/scaler

push: 
		docker push tarunpothulapati/scaler

dep:
		kubectl apply -f ./deploy/manifests.yaml