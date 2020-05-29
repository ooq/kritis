# Signer integration testing script
set -ex

# build a "good" example image
GOOD_IMAGE_URL=gcr.io/$PROJECT_ID/signer-int-good-image:$BUILD_ID

docker build -t $GOOD_IMAGE_URL -f ./Dockerfile.good .
delete_good_image () {
    ARG=$?
    echo "Delete good image."
    gcloud container images delete $GOOD_IMAGE_URL --force-delete-tags \
      --quiet
    exit $ARG
}
trap delete_good_image EXIT

docker push $GOOD_IMAGE_URL
# get image url with digest format
GOOD_IMG_DIGEST_URL=$(docker image inspect $GOOD_IMAGE_URL --format '{{index .RepoDigests 0}}')

./wait-for-analysis.py $GOOD_IMG_DIGEST_URL $PROJECT_ID

# build a "bad" example image


# check project id
echo $PROJECT_ID
echo $BUILD_ID