package advanced

import (
	"cmd-scaffold/log"
	"cmd-scaffold/pkg/em/direct"
)

func IsPackageDeployed(rb *direct.RequestBuilder, emClusterId int, productName, pkgVersion string) (bool, error) {
	// Get deployed product
	reqGetDeployed := &direct.RequestGetDeployedProduct{
		EmClusterId: emClusterId,
		ProductName: productName,
	}
	resGetDeployed, err := rb.GetDeployedProduct(reqGetDeployed)
	if err != nil {
		return false, nil
	}

	if resGetDeployed.Count == 0 || resGetDeployed.List[0].ProductVersion != pkgVersion {
		log.Info("product '%s' is not deployed or the version is different", productName)
		return false, nil
	}

	return true, nil
}
