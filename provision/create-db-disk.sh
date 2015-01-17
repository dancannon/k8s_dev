gcloud compute disks create --size="20GB" --zone=europe-west1-c rethinkdb-disk
gcloud compute instances attach-disk --zone=europe-west1-c --disk=rethinkdb-disk --device-name temp-data k8s-gke-master
gcloud compute ssh --zone=europe-west1-c k8s-gke-master \
  --command "sudo rm -rf /mnt/tmp && sudo mkdir /mnt/tmp && sudo /usr/share/google/safe_format_and_mount /dev/disk/by-id/google-temp-data /mnt/tmp"
gcloud compute instances detach-disk --zone=europe-west1-c --disk rethinkdb-disk k8s-gke-master

gcloud compute disks create --size="20GB" --zone=europe-west1-c consul-disk
gcloud compute instances attach-disk --zone=europe-west1-c --disk=consul-disk --device-name temp-data k8s-gke-master
gcloud compute ssh --zone=europe-west1-c k8s-gke-master \
  --command "sudo rm -rf /mnt/tmp && sudo mkdir /mnt/tmp && sudo /usr/share/google/safe_format_and_mount /dev/disk/by-id/google-temp-data /mnt/tmp"
gcloud compute instances detach-disk --zone=europe-west1-c --disk consul-disk k8s-gke-master

