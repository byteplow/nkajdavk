# nakjdavk
nak is hacked to gethere tool, to collect Notfall Kontakte (Emergency contacts) for our youth groups.

## install
Create secret for basic auth:
```
kubectl create secret generic RELEASENAME-nkajdavk --from-literal=user=USER --from-literal=password=PASSWORD
```

Install with helm:
```
helm upgrade --install RELEASENAME oci://registry-1.docker.io/byteplow/nkajdavk --version TAG
```
For current tag see: https://hub.docker.com/r/byteplow/nkajdavk/tags

## usage
Share the link `https://DOMAIN` with the partictipants. Let them fill out the form.

You can then see all submitted contacts. Therefor go to `https://USER:PASSWORD@DOMAIN/list`