# Paperless Office

<img src="assets/leaf.svg" height="100px">

The future is paperless. Unfortunately, most authorities (at least in Germany) still love paper, rendering 'digitalization' a foreign word.

To bypass this stale state, I started this project which allows scanned (and therefore digitized) documents to be...
* ... archived and stored savely
* ... (full-text) searched
* ... tagged

The goal for this project was not only environmental happiness, but also customer happiness.

Some of the key-features are:

* Automatic OCR of scanned documents enables full-text search on any document.
* Manual tagging improves searching the right document at the right time.
* Automatic Git backups reduce waste of physical space when compared to paper backups.
* An optional web viewer allows access to your documents. Anywhere, anytime.

## Prerequisites

* A (document) scanner (necessary)
* A container engine like Docker (recommended)
* docker-compose (recommended)
* Access to any Git provider (e.g. [GitLab](https://gitlab.com) or [GitHub](https://github.com), recommeded)

## Screenshots

![](assets/screenshots.png)

## How it works

1. Scan a document to a PDF file
2. Either upload the scanned PDF via the paperless-office Web UI or just dump it into the raw documents folder (see TODO: Configuration)
3. Let the server side do the text recognition magic. Once finished, the new document is accessible via the `Unconfirmed` section via the Web UI
4. Add some tags, double-check dates and recognized meta data (like URLs, e-mail addressess...)
5. Save and confirm the document. Saving triggers pushing the file to the Git repository (if configured). From now on, the document is prepared to be found in your paperless office.

## Setup

Using docker-compose is the most simple way to set up paperless-office.

* Prepare your environment:
  ```bash
  cd /path/paperless-office-documents
  mkdir -p data/raw data/processed

  # BEGIN Optionally init git repository
  cd data/processed
  git init

  # At the moment only https basic auth for Git is supported
  git remote add origin https://username:password@your.git/username/paperless-office-documents.git

  # Create .gitlab-ci.yaml or GitHub actions workflow, depending on your Git provider. For a GitLab snippet see further below.
  touch .gitlab-ci.yaml

  # Get the webviewer.json and index.html
  wget https://github.com/swinkelhofer/paperless-office/releases/latest/download/webviewer.js
  wget https://github.com/swinkelhofer/paperless-office/releases/latest/download/index.html

  # Stage and push your initial changes
  git add -A
  git commit -am "Init"
  git push

  cd ../..
  # END Optionally init git repository

  # Display your user ID for configuration in the next step
  id -u
  ```
* Save and adjust the following snippet to a file named `docker-compose.yaml`:
  ```yaml
  version: "3.6"
  services:
    paperless-office:
      image: ghcr.io/swinkelhofer/paperless-office:latest
      # user must match the UID of the volumes' owner
      user: "1000:1000"
      ports:
        - "8000:8000"
      volumes:
        - /path/paperless-office-documents/data/processed:/srv/data/processed
        - /path/paperless-office-documents/raw:/srv/data/raw
      restart: always
  ```
* Run `docker-compose up -d` to start paperless-office.
* Configured with the snippet above, the Web UI will be available via browser on http://localhost:8000/.


# Similar Projects
There are two other projects named [Paperless](https://github.com/the-paperless-project/paperless) and [Mayan EDMS](https://www.mayan-edms.com/) out there, that have technical overlap with paperless-office. In contrast to paperless-office, both are written in Python and do have a broader feature set (like document encryption). In favor, paperless-office brings a prettier UI, Git integration and a Webviewer allowing access to your documents via GitLab or GitHub pages.

# Git integration

A simple Git integration can be extended by supplying a CI workflow to deploy the contents via GitLab Pages or GitHub Pages.

## GitLab CI

`gitlab-ci.yaml` example configuration:
```yaml
pages:
  image: alpine:3.13
  script:
    - mkdir public
    - cp -rf * public/ || true
  artifacts:
    paths:
      - public
```

Each operation in paperless-office Web UI leads to a push to your Git repository. The CI pipeline will be triggered on each push, therefore re-deploying GitLab pages. The webviewer is then available via https://username.gitlab.io/paperless-office-documents

# Contribution

See [the contribution guidelines](CONTRIBUTING.md)