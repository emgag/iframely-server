# iframely-server

[![build](https://github.com/emgag/iframely-server/actions/workflows/build.yml/badge.svg)](https://github.com/emgag/iframely-server/actions/workflows/build.yml)
[![release](https://github.com/emgag/iframely-server/actions/workflows/release.yml/badge.svg)](https://github.com/emgag/iframely-server/actions/workflows/release.yml)

Simple wrapper for iframely oEmbed API.

## Endpoints

### _/info_

**Parameters**:

- (required) _api_key_: The API key for authentication
- (required) _url_: The url of the requested embed

**Response**:

A JSON object with the following values:

- _provider_: The lowercase embed provider name
- _url_: The requested embed url
- _html_: The embed html code
- _image_: The embed (preview) image

or 

- _error_: In case of an error returned by iframely

## Environment variables

- `SERVER_API_KEY`: Key to authenticate clients
- `IFRAMELY_API_KEY`: API key for fetching data from iframely
