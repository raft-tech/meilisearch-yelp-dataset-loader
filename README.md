# Meilisearch Yelp Dataset Loader

### Purpose

This application can be used to test [Meilisearch](https://www.meilisearch.com/docs) with some datasets provided by
Yelp. See [here](https://www.yelp.com/dataset) for more information.

### Running

Due to the size of the datasets, they exceeded GitHub's file sizer limit, so they will first need to be downloaded from
[here](https://www.yelp.com/dataset/download). Once downloaded and extracted, move them to the repo data directory and
update main.go with the filenames of the datasets to upload.

To run the application, first ensure Meilisearch is running, using whatever is preferred. You can execute
a `docker compose up -d` to run a Meilisearch server locally running in detached mode.

Update the host in main.go to point to the Meilisearch host, and then execute `make run` from the command line to load
the Yelp datasets into Meilisearch.