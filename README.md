# Description
K8s operator that Monitor the SSL certificates on K8s that are not managed by cert-manager like GoDaddy certifications and send alerts before certificate expiration 

# Feature
- Containous monitor certificates in the cluster
- Define alert annotations for certificates and provide diffrent alert mechansim
- Alert methods
  - calling http API
  - Microsoft team channel 
  - Slack
- Auto re-new certifiates by watching ingress and automatic pull the new TLS from 
  - custom file server via download cert and key file using http 
  - s3
  - min.io
  
