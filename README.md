# Meilisearch Yelp Dataset Loader

### Purpose

This application can be used to test [Meilisearch](https://www.meilisearch.com/docs) with some datasets provided by
Yelp. See [here](https://www.yelp.com/dataset) for more information.

### Running

Due to the large size of the datasets, they exceed GitHub's file size limit, so they will first need to be downloaded
from [here](https://www.yelp.com/dataset/download). Once downloaded and extracted, move them to the repository data
directory and update main.go with the indexes, filenames, and primary keys of the datasets to upload.

To run the application, first, ensure an instance of Meilisearch is running. You can execute
a `docker compose up -d` to run a Meilisearch server locally in detached mode.

Update the host in main.go to point to the Meilisearch host, and then execute `make run` from the command line to load
the Yelp datasets into Meilisearch.
