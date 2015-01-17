#! /bin/bash
gcloud compute ssh k8s-gke-node-1 --ssh-flag="-L 5000:10.147.250.201:5000" --ssh-flag="-N" --ssh-flag="-n"
