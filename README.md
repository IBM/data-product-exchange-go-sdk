# Data Product Hub Go SDK Version 3.0.0
Go client library to interact with the various Data Product Hub Service APIs

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  * [Go modules](#go-modules)
  * [`go get` command](#go-get-command)
- [Using the SDK](#using-the-sdk)
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop --> 

## Overview

The IBM Cloud Data Product Hub Go SDK allows developers to programmatically interact with the following services:

Service Name | Package name 
--- | --- 
[Data Product Hub](https://cloud.ibm.com/apidocs/dataproducts) | dataproducthubapiservicev1


## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.19 or above.

## Installation
The current version of this SDK: 3.0.0

### Go modules  
If your application uses Go modules for dependency management (recommended), just add an import for the service 
that you will use in your application.  
Here is an example:

```go
import (
	"github.com/IBM/data-product-exchange-go-sdk/dataproducthubapiservicev1"
)
```
Next, run `go build` or `go mod tidy` to download and install the new dependencies and update your application's
`go.mod` file.  

In the example above, the `dataproducthubapiservicev1` part of the import path is the package name
associated with the DPH service..

### `go get` command  
Alternatively, you can use the `go get` command to download and install the appropriate packages needed by your application:
```
go get -u github.com/IBM/data-product-exchange-go-sdk/dataproducthubapiservicev1
```

## Using the SDK
For general SDK usage information, please see [this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/main/README.md)

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at 
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](https://github.com/IBM/data-product-exchange-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/).

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
