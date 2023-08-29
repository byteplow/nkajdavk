build:
	podman build -t docker.io/byteplow/nkajdavk .

helmupdate:
	helm dependency update contrib/deployment/chart

helmbuild:
	sed -i "s/gitshaplaceholder/$$(git rev-parse --short HEAD)/g" contrib/deployment/chart/values.yaml
	sed -i "s/gitshaplaceholder/$$(git rev-parse --short HEAD)/g" contrib/deployment/chart/Chart.yaml
	helm dependency build contrib/deployment/chart
	helm package contrib/deployment/chart
	sed -i "s/$$(git rev-parse --short HEAD)/gitshaplaceholder/g" contrib/deployment/chart/values.yaml
	sed -i "s/$$(git rev-parse --short HEAD)/gitshaplaceholder/g" contrib/deployment/chart/Chart.yaml

publisch: build helmbuild
	podman tag docker.io/byteplow/nkajdavk docker.io/byteplow/nkajdavk:$$(git rev-parse --short HEAD)
	podman push docker.io/byteplow/nkajdavk:$$(git rev-parse --short HEAD)

	helm push nkajdavk-*-$$(git rev-parse --short HEAD).tgz oci://docker.io/byteplow