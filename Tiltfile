allow_k8s_contexts('gke_first-torus-238116_us-central1-c_new-cluster')

# Deploy: tell Tilt what YAML to deploy
k8s_yaml('./manifests/deployment.yaml')

# Build: tell Tilt what images to build from which directories
docker_build('tarunpothulapati/scaler', './')
