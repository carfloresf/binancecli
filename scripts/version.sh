MAJOR_VERSION=0
MINOR_VERSION=1
BASE_VERSION="${MAJOR_VERSION}.${MINOR_VERSION}-**PATCH**"
PATCH_VERSION=$(git log -1 --date=format:%Y%m%d.%H%M%S --pretty=tformat:%cd-g%h)
VERSION=`echo -n ${BASE_VERSION} | sed -e "s/\*\*PATCH\*\*/${PATCH_VERSION}/"`
echo $VERSION
