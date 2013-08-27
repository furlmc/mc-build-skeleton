set -e
shopt -s nullglob

source src/build.txt

./modid > target_info
source target_info
rm target_info

echo "${TARGET_NAME}: Starting build"

cd forge/mcp

# Copy in source from git
cp -r -t src/minecraft ../../src/$SOURCE

./recompile.sh

./reobfuscate.sh

cd ../..

if [ ! -e build_number ]; then
	echo 0 > build_number
fi

read BUILD_NUMBER < build_number
BUILD_NUMBER=$[$BUILD_NUMBER + 1]
echo $BUILD_NUMBER > build_number

source src/build.txt

./jarbuilder -assets="${ASSETS}" -v="${VERSION}" -mc="${MCVERSION}" -ns="${SOURCE}" -filename="dist/${TARGET}-${MCVERSION}-${VERSION}.zip"

echo "${TARGET_NAME}: Build completed"
