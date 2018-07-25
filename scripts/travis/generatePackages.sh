#!/bin/bash

set -x

setup_fpm() {
	if which fpm >/dev/null; then
		echo "FPM already exists"
	else
		echo "Installing FPM... "
		gem install fpm
	fi
}

setup_fpm

# Remove all posible old packages
rm -rf ./packages

for package_type in ${PACKAGES_TYPES}; do
	mkdir -p ./packages/${package_type}
	fpm -s dir -t ${package_type} -n panda \
		-p ./packages/${package_type}  \
		--version "${PANDA_VERSION}" \
		--log error \
		--vendor="smmx" \
		--url="www.segundamano.mx" \
		--provides="panda" \
		--description="Handle orders for Segundamano" \
		--iteration=${TRAVIS_BUILD_NUMBER} \
		-C ./build \
		opt/panda/bin/panda
done
