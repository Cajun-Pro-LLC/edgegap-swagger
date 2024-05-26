# Go API client for swagger

<h1 style=\"margin-top: 2rem\">Introduction</h1> <p>The Edgegap API lets you manage all your resources from standard HTTP requests. We promote using them to automate all your processes that are using Edgegap.</p> <p>If you have any questions, don't hesitate to contact us via email, or you can also jump on our <a href=\"https://discord.com/invite/GYaHcKR9a5\" target=\"_blank\">Discord</a>. We will be happy to help. Feel free to make features request; we also love those.</p> <div class=\"theme-doc-markdown markdown\"><h2 class=\"anchor anchorWithStickyNavbar_node_modules-@docusaurus-theme-classic-lib-theme-Heading-styles-module\" id=\"pagination---response\">Pagination - Response<a class=\"hash-link\" href=\"#pagination---response\" title=\"Direct link to heading\">​</a></h2><p>The GET response body can be returned with pagination to avoid requesting all the data at once.</p><p>Pagination object will always be under the  <em>paginator</em> key.</p><div class=\"language-json codeBlockContainer_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Container-styles-module theme-code-block\" style=\"--prism-color:#bfc7d5; --prism-background-color:#292d3e;\"><div class=\"codeBlockContent_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><pre tabindex=\"0\" class=\"prism-code language-json codeBlock_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module thin-scrollbar\"><code class=\"codeBlockLines_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">{</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"number\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"The Current page, default=1\"</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"next_page_number\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"The Next page number or null\"</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"previous_page_number\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"The Previous page number or null\"</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"paginator\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">{</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">    </span><span class=\"token property\">\"num_pages\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"The total numbers of pages\"</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">}</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"has_next\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"Boolean if there is a next page\"</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"has_previous\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"Boolean if there is a previous page\"</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\"></span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">}</span><br></span></code></pre><div class=\"buttonGroup_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><button type=\"button\" aria-label=\"Copy code to clipboard\" title=\"Copy\" class=\"clean-btn\"><span class=\"copyButtonIcons_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" aria-hidden=\"true\"><svg class=\"copyButtonIcon_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" viewBox=\"0 0 24 24\"><path d=\"M19,21H8V7H19M19,5H8A2,2 0 0,0 6,7V21A2,2 0 0,0 8,23H19A2,2 0 0,0 21,21V7A2,2 0 0,0 19,5M16,1H4A2,2 0 0,0 2,3V17H4V3H16V1Z\"></path></svg><svg class=\"copyButtonSuccessIcon_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" viewBox=\"0 0 24 24\"><path d=\"M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z\"></path></svg></span></button></div></div></div><p>Full Body Example:</p><div class=\"language-json codeBlockContainer_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Container-styles-module theme-code-block\" style=\"--prism-color:#bfc7d5; --prism-background-color:#292d3e;\"><div class=\"codeBlockContent_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><pre tabindex=\"0\" class=\"prism-code language-json codeBlock_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module thin-scrollbar\"><code class=\"codeBlockLines_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">{</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"count\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token number\" style=\"color: rgb(247, 140, 108);\">100</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"data\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">[</span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"value-0\"</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"> </span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"[...]\"</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"> </span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"value-9\"</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">]</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"success\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token boolean\" style=\"color: rgb(255, 88, 116);\">true</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"pagination\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">{</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">    </span><span class=\"token property\">\"number\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token number\" style=\"color: rgb(247, 140, 108);\">1</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">    </span><span class=\"token property\">\"next_page_number\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token number\" style=\"color: rgb(247, 140, 108);\">2</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">    </span><span class=\"token property\">\"previous_page_number\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token null keyword\" style=\"font-style: italic;\">null</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">    </span><span class=\"token property\">\"paginator\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">{</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">      </span><span class=\"token property\">\"num_pages\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token number\" style=\"color: rgb(247, 140, 108);\">10</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">    </span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">}</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">    </span><span class=\"token property\">\"has_next\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token boolean\" style=\"color: rgb(255, 88, 116);\">true</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">    </span><span class=\"token property\">\"has_previous\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token boolean\" style=\"color: rgb(255, 88, 116);\">false</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">}</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\"></span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">}</span><br></span></code></pre><div class=\"buttonGroup_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><button type=\"button\" aria-label=\"Copy code to clipboard\" title=\"Copy\" class=\"clean-btn\"><span class=\"copyButtonIcons_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" aria-hidden=\"true\"><svg class=\"copyButtonIcon_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" viewBox=\"0 0 24 24\"><path d=\"M19,21H8V7H19M19,5H8A2,2 0 0,0 6,7V21A2,2 0 0,0 8,23H19A2,2 0 0,0 21,21V7A2,2 0 0,0 19,5M16,1H4A2,2 0 0,0 2,3V17H4V3H16V1Z\"></path></svg><svg class=\"copyButtonSuccessIcon_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" viewBox=\"0 0 24 24\"><path d=\"M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z\"></path></svg></span></button></div></div></div><h2 class=\"anchor anchorWithStickyNavbar_node_modules-@docusaurus-theme-classic-lib-theme-Heading-styles-module\" id=\"pagination---parameters\">Pagination - Parameters<a class=\"hash-link\" href=\"#pagination---parameters\" title=\"Direct link to heading\">​</a></h2><p>You can add those values to manipulate the pagination object in the URL Parameters.</p><ul class=\"\"><li>page</li><li>limit</li></ul><p>Example:</p><div class=\"language-text codeBlockContainer_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Container-styles-module theme-code-block\" style=\"--prism-color:#bfc7d5; --prism-background-color:#292d3e;\"><div class=\"codeBlockContent_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><pre tabindex=\"0\" class=\"prism-code language-text codeBlock_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module thin-scrollbar\"><code class=\"codeBlockLines_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\"># To get the second page</span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">GET - https://api.edgegap.com/v1/apps?page=2</span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\" style=\"display: inline-block;\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\"># To change the count of element in one page (20/page)</span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">GET - https://api.edgegap.com/v1/apps?limit=20</span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\" style=\"display: inline-block;\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\"># You can mix those (20/page, second page)</span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">GET - https://api.edgegap.com/v1/apps?page=2&amp;limit=20</span><br></span></code></pre><div class=\"buttonGroup_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><button type=\"button\" aria-label=\"Copy code to clipboard\" title=\"Copy\" class=\"clean-btn\"><span class=\"copyButtonIcons_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" aria-hidden=\"true\"><svg class=\"copyButtonIcon_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" viewBox=\"0 0 24 24\"><path d=\"M19,21H8V7H19M19,5H8A2,2 0 0,0 6,7V21A2,2 0 0,0 8,23H19A2,2 0 0,0 21,21V7A2,2 0 0,0 19,5M16,1H4A2,2 0 0,0 2,3V17H4V3H16V1Z\"></path></svg><svg class=\"copyButtonSuccessIcon_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" viewBox=\"0 0 24 24\"><path d=\"M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z\"></path></svg></span></button></div></div></div></div>

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen
For more information, please visit [https://edgegap.com](https://edgegap.com)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "github.com/cajun-pro-llc/edgegap-swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApplicationsApi* | [**AppVersionDelete**](docs/ApplicationsApi.md#appversiondelete) | **Delete** /v1/app/{app_name}/version/{version_name} | Delete an Application Version
*ApplicationsApi* | [**AppVersionGet**](docs/ApplicationsApi.md#appversionget) | **Get** /v1/app/{app_name}/version/{version_name} | Get an Application Version
*ApplicationsApi* | [**AppVersionPost**](docs/ApplicationsApi.md#appversionpost) | **Post** /v1/app/{app_name}/version | Create a New Application Version
*ApplicationsApi* | [**AppVersionWhitelistEntryDelete**](docs/ApplicationsApi.md#appversionwhitelistentrydelete) | **Delete** /v1/app/{app_name}/version/{version_name}/whitelist/{entry_id} | Delete an ACL Entry
*ApplicationsApi* | [**AppVersionWhitelistEntryGet**](docs/ApplicationsApi.md#appversionwhitelistentryget) | **Get** /v1/app/{app_name}/version/{version_name}/whitelist/{entry_id} | Get an ACL Entry
*ApplicationsApi* | [**AppVersionWhitelistGet**](docs/ApplicationsApi.md#appversionwhitelistget) | **Get** /v1/app/{app_name}/version/{version_name}/whitelist | List All ACL Entries for an Application Version
*ApplicationsApi* | [**AppVersionWhitelistPost**](docs/ApplicationsApi.md#appversionwhitelistpost) | **Post** /v1/app/{app_name}/version/{version_name}/whitelist | Create an ACL Entry
*ApplicationsApi* | [**AppVersionsGet**](docs/ApplicationsApi.md#appversionsget) | **Get** /v1/app/{app_name}/versions | List All Versions for an Application
*ApplicationsApi* | [**AppVersionsPatch**](docs/ApplicationsApi.md#appversionspatch) | **Patch** /v1/app/{app_name}/version/{version_name} | Update an Application Version
*ApplicationsApi* | [**ApplicationDelete**](docs/ApplicationsApi.md#applicationdelete) | **Delete** /v1/app/{app_name} | Delete an Application
*ApplicationsApi* | [**ApplicationGet**](docs/ApplicationsApi.md#applicationget) | **Get** /v1/app/{app_name} | Get an Application
*ApplicationsApi* | [**ApplicationPatch**](docs/ApplicationsApi.md#applicationpatch) | **Patch** /v1/app/{app_name} | Update an Application
*ApplicationsApi* | [**ApplicationPost**](docs/ApplicationsApi.md#applicationpost) | **Post** /v1/app | Create a New Application
*ApplicationsApi* | [**ApplicationsGet**](docs/ApplicationsApi.md#applicationsget) | **Get** /v1/apps | List All Applications
*ContextApi* | [**ContextCreateDeploymentTag**](docs/ContextApi.md#contextcreatedeploymenttag) | **Post** /v1/context/{request_id}/{security_number}/tags | Create a Tag for a Running Deployment
*ContextApi* | [**ContextDeleteDeploymentTag**](docs/ContextApi.md#contextdeletedeploymenttag) | **Delete** /v1/context/{request_id}/{security_number}/tags/{tag_name} | Delete a Tag from a Running Deployment
*ContextApi* | [**ContextGet**](docs/ContextApi.md#contextget) | **Get** /v1/context/{request_id}/{security_number} | Get the Context of a Deployment
*CustomSessionsApi* | [**DeleteCustomSession**](docs/CustomSessionsApi.md#deletecustomsession) | **Delete** /v1/deployment/{request_id}/custom/session/{custom_id} | Delete a Custom Session
*CustomSessionsApi* | [**DeleteCustomSessions**](docs/CustomSessionsApi.md#deletecustomsessions) | **Delete** /v1/deployment/{request_id}/custom/sessions | Delete Custom Sessions in Bulk
*CustomSessionsApi* | [**GetCustomSession**](docs/CustomSessionsApi.md#getcustomsession) | **Get** /v1/deployment/{request_id}/custom/session/{custom_id} | Get a Custom Session
*CustomSessionsApi* | [**GetCustomSessions**](docs/CustomSessionsApi.md#getcustomsessions) | **Get** /v1/deployment/{request_id}/custom/sessions | List All Custom Session of a Deployment
*CustomSessionsApi* | [**PatchCustomSession**](docs/CustomSessionsApi.md#patchcustomsession) | **Patch** /v1/deployment/{request_id}/custom/session/{custom_id} | Update a Custom Session
*CustomSessionsApi* | [**PostCustomSession**](docs/CustomSessionsApi.md#postcustomsession) | **Post** /v1/deployment/{request_id}/custom/session/{custom_id} | Create a New Custom Session
*CustomSessionsApi* | [**PostCustomSessions**](docs/CustomSessionsApi.md#postcustomsessions) | **Post** /v1/deployment/{request_id}/custom/sessions | Create Custom Sessions in Bulk
*DeploymentTagsApi* | [**DeploymentsTagsCreate**](docs/DeploymentTagsApi.md#deploymentstagscreate) | **Post** /v1/deployments/{request_id}/tags | Create Tag for a Deployment
*DeploymentTagsApi* | [**DeploymentsTagsDelete**](docs/DeploymentTagsApi.md#deploymentstagsdelete) | **Delete** /v1/deployments/{request_id}/tags/{tag_name} | Delete Tag for a Deployment
*DeploymentTagsApi* | [**DeploymentsTagsList**](docs/DeploymentTagsApi.md#deploymentstagslist) | **Get** /v1/deployments/{request_id}/tags | List tags for a Deployment
*DeploymentTagsApi* | [**DeploymentsTagsRead**](docs/DeploymentTagsApi.md#deploymentstagsread) | **Get** /v1/deployments/{request_id}/tags/{tag_name} | Get tag for a Deployment
*DeploymentTagsApi* | [**DeploymentsTagsUpdate**](docs/DeploymentTagsApi.md#deploymentstagsupdate) | **Patch** /v1/deployments/{request_id}/tags/{tag_name} | Update Tag for a Deployment
*DeploymentsApi* | [**Deploy**](docs/DeploymentsApi.md#deploy) | **Post** /v1/deploy | Create a Deployment
*DeploymentsApi* | [**DeploymentDelete**](docs/DeploymentsApi.md#deploymentdelete) | **Delete** /v1/stop/{request_id} | Delete a Deployment
*DeploymentsApi* | [**DeploymentGetLogs**](docs/DeploymentsApi.md#deploymentgetlogs) | **Get** /v1/deployment/{request_id}/container-logs | Get Deployment Container Logs
*DeploymentsApi* | [**DeploymentStatusGet**](docs/DeploymentsApi.md#deploymentstatusget) | **Get** /v1/status/{request_id} | Get a Deployment Status and Information
*DeploymentsApi* | [**DeploymentUpdate**](docs/DeploymentsApi.md#deploymentupdate) | **Patch** /v1/deployments/{request_id} | Updates properties of a deployment
*DeploymentsApi* | [**DeploymentsBulkDelete**](docs/DeploymentsApi.md#deploymentsbulkdelete) | **Post** /v1/deployments/bulk-stop | Delete Deployments in Bulk
*DeploymentsApi* | [**DeploymentsGet**](docs/DeploymentsApi.md#deploymentsget) | **Get** /v1/deployments | List All Deployments
*DeploymentsApi* | [**SelfDeploymentDelete**](docs/DeploymentsApi.md#selfdeploymentdelete) | **Delete** /v1/self/stop/{request_id}/{access_point_id} | Delete a Deployment from inside the container
*EndpointStorageApi* | [**EndpointCreate**](docs/EndpointStorageApi.md#endpointcreate) | **Post** /v1/storage/endpoint | Create a New Endpoint Storage
*EndpointStorageApi* | [**EndpointDelete**](docs/EndpointStorageApi.md#endpointdelete) | **Delete** /v1/storage/endpoint/{endpoint_name} | Delete an Endpoint Storage
*EndpointStorageApi* | [**EndpointGet**](docs/EndpointStorageApi.md#endpointget) | **Get** /v1/storage/endpoint/{endpoint_name} | Get an Endpoint Storage
*EndpointStorageApi* | [**EndpointUpdate**](docs/EndpointStorageApi.md#endpointupdate) | **Patch** /v1/storage/endpoint/{endpoint_name} | Update an Endpoint Storage
*EndpointStorageApi* | [**EndpointsList**](docs/EndpointStorageApi.md#endpointslist) | **Get** /v1/storage/endpoints | List All Endpoint Storage
*EndpointStorageApi* | [**PullProfileCreate**](docs/EndpointStorageApi.md#pullprofilecreate) | **Post** /v1/storage/endpoint/{endpoint_name}/pull-profile | Create a New Pull Profile
*EndpointStorageApi* | [**PullProfileDelete**](docs/EndpointStorageApi.md#pullprofiledelete) | **Delete** /v1/storage/endpoint/{endpoint_name}/pull-profile/{pull_profile_name} | Delete a Pull Profile
*EndpointStorageApi* | [**PullProfileGet**](docs/EndpointStorageApi.md#pullprofileget) | **Get** /v1/storage/endpoint/{endpoint_name}/pull-profile/{pull_profile_name} | Get a Pull Profile
*EndpointStorageApi* | [**PullProfileLinkAppVersion**](docs/EndpointStorageApi.md#pullprofilelinkappversion) | **Put** /v1/storage/endpoint/{endpoint_name}/pull-profile/{pull_profile_name}/app/{app_name}/version/{version_name} | Link a Pull Profile to an Application Version
*EndpointStorageApi* | [**PullProfileList**](docs/EndpointStorageApi.md#pullprofilelist) | **Get** /v1/storage/endpoint/{endpoint_name}/pull-profiles | List All Pull Profile of an Endpoint Storage
*EndpointStorageApi* | [**PullProfileUnlinkAppVersion**](docs/EndpointStorageApi.md#pullprofileunlinkappversion) | **Delete** /v1/storage/endpoint/{endpoint_name}/pull-profile/{pull_profile_name}/app/{app_name}/version/{version_name} | Unlink a Pull Profile From an Application Version
*EndpointStorageApi* | [**PullProfileUpdate**](docs/EndpointStorageApi.md#pullprofileupdate) | **Patch** /v1/storage/endpoint/{endpoint_name}/pull-profile/{pull_profile_name} | Update a Pull Profile
*FleetsApi* | [**FleetCreate**](docs/FleetsApi.md#fleetcreate) | **Post** /v1/fleet | Create a Fleet
*FleetsApi* | [**FleetDelete**](docs/FleetsApi.md#fleetdelete) | **Delete** /v1/fleet/{fleet_name} | Delete a Fleet
*FleetsApi* | [**FleetGet**](docs/FleetsApi.md#fleetget) | **Get** /v1/fleet/{fleet_name} | Get a Fleet
*FleetsApi* | [**FleetLinkAppVersion**](docs/FleetsApi.md#fleetlinkappversion) | **Put** /v1/fleet/{fleet_name}/app/{app_name}/version/{version_name} | Link an Application Version to a Fleet
*FleetsApi* | [**FleetPoliciesCreate**](docs/FleetsApi.md#fleetpoliciescreate) | **Post** /v1/fleet/{fleet_name}/policies | Create a Fleet Policy
*FleetsApi* | [**FleetPoliciesDelete**](docs/FleetsApi.md#fleetpoliciesdelete) | **Delete** /v1/fleet/{fleet_name}/policies/{policy_name} | Delete a Policy
*FleetsApi* | [**FleetPoliciesGet**](docs/FleetsApi.md#fleetpoliciesget) | **Get** /v1/fleet/{fleet_name}/policies/{policy_name} | Get a Policy
*FleetsApi* | [**FleetPoliciesList**](docs/FleetsApi.md#fleetpolicieslist) | **Get** /v1/fleet/{fleet_name}/policies | List All Policies of a Fleet
*FleetsApi* | [**FleetPoliciesUpdate**](docs/FleetsApi.md#fleetpoliciesupdate) | **Patch** /v1/fleet/{fleet_name}/policies/{policy_name} | Update a Policy
*FleetsApi* | [**FleetUnlinkAppVersion**](docs/FleetsApi.md#fleetunlinkappversion) | **Delete** /v1/fleet/{fleet_name}/app/{app_name}/version/{version_name} | Unlink an Application Version From a Fleet
*FleetsApi* | [**FleetUpdate**](docs/FleetsApi.md#fleetupdate) | **Patch** /v1/fleet/{fleet_name} | Update a Fleet
*FleetsApi* | [**Fleets**](docs/FleetsApi.md#fleets) | **Get** /v1/fleets | List All Fleets
*IPLookupApi* | [**IP**](docs/IPLookupApi.md#ip) | **Get** /v1/ip | Get Your Public IP
*IPLookupApi* | [**IPLookup**](docs/IPLookupApi.md#iplookup) | **Get** /v1/ip/{ip}/lookup | Get an IP&#39;s information
*IPLookupApi* | [**IPsLookup**](docs/IPLookupApi.md#ipslookup) | **Post** /v1/ips/lookup | Get IPs Information in Bulk
*LobbiesApi* | [**LobbyCreate**](docs/LobbiesApi.md#lobbycreate) | **Post** /v1/lobbies | Create a Lobby
*LobbiesApi* | [**LobbyDeploy**](docs/LobbiesApi.md#lobbydeploy) | **Post** /v1/lobbies:deploy | Deploy a Lobby
*LobbiesApi* | [**LobbyGet**](docs/LobbiesApi.md#lobbyget) | **Get** /v1/lobbies/{lobby_name} | Get a Lobby
*LobbiesApi* | [**LobbyTerminate**](docs/LobbiesApi.md#lobbyterminate) | **Post** /v1/lobbies:terminate | Terminate a Lobby
*LocationsApi* | [**LocationBeaconList**](docs/LocationsApi.md#locationbeaconlist) | **Get** /v1/locations/beacons | List All Location Beacons
*LocationsApi* | [**LocationsGet**](docs/LocationsApi.md#locationsget) | **Get** /v1/locations | List All Locations
*MatchmakerApi* | [**DeleteMatchmaker**](docs/MatchmakerApi.md#deletematchmaker) | **Delete** /v1/aom/matchmaker/{matchmaker_name} | Delete a Matchmaker
*MatchmakerApi* | [**DeleteMatchmakerComponent**](docs/MatchmakerApi.md#deletematchmakercomponent) | **Delete** /v1/aom/component/{component_name} | Delete a Matchmaker Component
*MatchmakerApi* | [**DeleteMatchmakerComponentEnv**](docs/MatchmakerApi.md#deletematchmakercomponentenv) | **Delete** /v1/aom/component/{component_name}/env/{env_key} | Delete a Matchmaker Component ENV
*MatchmakerApi* | [**DeleteMatchmakerManagedRelease**](docs/MatchmakerApi.md#deletematchmakermanagedrelease) | **Delete** /v1/aom/matchmaker/{matchmaker_name}/release/managed/{release_version} | Delete a Matchmaker Managed Release
*MatchmakerApi* | [**DeleteMatchmakerRelease**](docs/MatchmakerApi.md#deletematchmakerrelease) | **Delete** /v1/aom/matchmaker/{matchmaker_name}/release/{release_version} | Delete a Matchmaker Release
*MatchmakerApi* | [**DeleteMatchmakerReleaseConfig**](docs/MatchmakerApi.md#deletematchmakerreleaseconfig) | **Delete** /v1/aom/release/config/{config_name} | Delete a Matchmaker Release Config
*MatchmakerApi* | [**GetComponentList**](docs/MatchmakerApi.md#getcomponentlist) | **Get** /v1/aom/components | List All Matchmaker Components
*MatchmakerApi* | [**GetEnvsList**](docs/MatchmakerApi.md#getenvslist) | **Get** /v1/aom/component/{component_name}/envs | List All Matchmaker Component ENVs
*MatchmakerApi* | [**GetMatchmaker**](docs/MatchmakerApi.md#getmatchmaker) | **Get** /v1/aom/matchmaker/{matchmaker_name} | Get a Matchmaker
*MatchmakerApi* | [**GetMatchmakerComponent**](docs/MatchmakerApi.md#getmatchmakercomponent) | **Get** /v1/aom/component/{component_name} | Get a Matchmaker Component
*MatchmakerApi* | [**GetMatchmakerComponentEnv**](docs/MatchmakerApi.md#getmatchmakercomponentenv) | **Get** /v1/aom/component/{component_name}/env/{env_key} | Get a Matchmaker Component ENV
*MatchmakerApi* | [**GetMatchmakerList**](docs/MatchmakerApi.md#getmatchmakerlist) | **Get** /v1/aom/matchmakers | List All Matchmakers
*MatchmakerApi* | [**GetMatchmakerManagedRelease**](docs/MatchmakerApi.md#getmatchmakermanagedrelease) | **Get** /v1/aom/matchmaker/{matchmaker_name}/release/managed/{release_version} | Get a Matchmaker Managed Release
*MatchmakerApi* | [**GetMatchmakerRelease**](docs/MatchmakerApi.md#getmatchmakerrelease) | **Get** /v1/aom/matchmaker/{matchmaker_name}/release/{release_version} | Get a Matchmaker Release
*MatchmakerApi* | [**GetMatchmakerReleaseConfig**](docs/MatchmakerApi.md#getmatchmakerreleaseconfig) | **Get** /v1/aom/release/config/{config_name} | Get a Matchmaker Release Config
*MatchmakerApi* | [**GetReleaseConfigsList**](docs/MatchmakerApi.md#getreleaseconfigslist) | **Get** /v1/aom/release/configs | List All Matchmaker Release Configs
*MatchmakerApi* | [**GetReleaseList**](docs/MatchmakerApi.md#getreleaselist) | **Get** /v1/aom/matchmaker/{matchmaker_name}/releases | List All Matchmaker Releases
*MatchmakerApi* | [**PatchMatchmaker**](docs/MatchmakerApi.md#patchmatchmaker) | **Patch** /v1/aom/matchmaker/{matchmaker_name} | Update a Matchmaker
*MatchmakerApi* | [**PatchMatchmakerComponent**](docs/MatchmakerApi.md#patchmatchmakercomponent) | **Patch** /v1/aom/component/{component_name} | Update a Matchmaker Component
*MatchmakerApi* | [**PatchMatchmakerComponentEnv**](docs/MatchmakerApi.md#patchmatchmakercomponentenv) | **Patch** /v1/aom/component/{component_name}/env/{env_key} | Update a Matchmaker Component ENV
*MatchmakerApi* | [**PatchMatchmakerManagedRelease**](docs/MatchmakerApi.md#patchmatchmakermanagedrelease) | **Patch** /v1/aom/matchmaker/{matchmaker_name}/release/managed/{release_version} | Update a Matchmaker Managed Release
*MatchmakerApi* | [**PatchMatchmakerRelease**](docs/MatchmakerApi.md#patchmatchmakerrelease) | **Patch** /v1/aom/matchmaker/{matchmaker_name}/release/{release_version} | Update a Matchmaker Release
*MatchmakerApi* | [**PatchMatchmakerReleaseConfig**](docs/MatchmakerApi.md#patchmatchmakerreleaseconfig) | **Patch** /v1/aom/release/config/{config_name} | Update a Matchmaker Release Config
*MatchmakerApi* | [**PostMatchmaker**](docs/MatchmakerApi.md#postmatchmaker) | **Post** /v1/aom/matchmaker | Create a Matchmaker
*MatchmakerApi* | [**PostMatchmakerComponent**](docs/MatchmakerApi.md#postmatchmakercomponent) | **Post** /v1/aom/component | Create a Matchmaker Component
*MatchmakerApi* | [**PostMatchmakerComponentEnv**](docs/MatchmakerApi.md#postmatchmakercomponentenv) | **Post** /v1/aom/component/{component_name}/env | Create a Matchmaker Component ENV
*MatchmakerApi* | [**PostMatchmakerManagedRelease**](docs/MatchmakerApi.md#postmatchmakermanagedrelease) | **Post** /v1/aom/matchmaker/{matchmaker_name}/release/managed | Create a Matchmaker Managed Release
*MatchmakerApi* | [**PostMatchmakerRelease**](docs/MatchmakerApi.md#postmatchmakerrelease) | **Post** /v1/aom/matchmaker/{matchmaker_name}/release | Create a Matchmaker Release
*MatchmakerApi* | [**PostMatchmakerReleaseConfig**](docs/MatchmakerApi.md#postmatchmakerreleaseconfig) | **Post** /v1/aom/release/config | Create a Matchmaker Release Config
*MetricsApi* | [**DeploymentMetricsGet**](docs/MetricsApi.md#deploymentmetricsget) | **Get** /v1/metrics/deployment/{request_id} | Get a Deployment Metrics
*MonitoringApi* | [**Monitoring**](docs/MonitoringApi.md#monitoring) | **Get** /monitor | 
*RelaysApi* | [**RelaySessionCreate**](docs/RelaysApi.md#relaysessioncreate) | **Post** /v1/relays/sessions | Create a Relay Session
*RelaysApi* | [**RelaySessionDelete**](docs/RelaysApi.md#relaysessiondelete) | **Delete** /v1/relays/sessions/{session_id} | Delete a Relay Session
*RelaysApi* | [**RelaySessionGet**](docs/RelaysApi.md#relaysessionget) | **Get** /v1/relays/sessions/{session_id} | Get a Relay Session
*RelaysApi* | [**RelaySessionList**](docs/RelaysApi.md#relaysessionlist) | **Get** /v1/relays/sessions | List all Relay Sessions
*RelaysApi* | [**RelayUserAuthorize**](docs/RelaysApi.md#relayuserauthorize) | **Post** /v1/relays/sessions:authorize-user | Authorize a user on a Relay Session
*RelaysApi* | [**RelayUserRevoke**](docs/RelaysApi.md#relayuserrevoke) | **Post** /v1/relays/sessions:revoke-user | Remove a user on a Relay Session
*SessionsApi* | [**DeleteUsersSession**](docs/SessionsApi.md#deleteuserssession) | **Delete** /v1/session/{session_id}/users | Delete Users From a Session
*SessionsApi* | [**GetSession**](docs/SessionsApi.md#getsession) | **Get** /v1/session/{session_id} | Get a Session
*SessionsApi* | [**GetUsersSession**](docs/SessionsApi.md#getuserssession) | **Get** /v1/session/{session_id}/users | List Users of a Session
*SessionsApi* | [**ListSessions**](docs/SessionsApi.md#listsessions) | **Get** /v1/sessions | List All Sessions
*SessionsApi* | [**PutUsersSession**](docs/SessionsApi.md#putuserssession) | **Put** /v1/session/{session_id}/users | Put Users in a Session
*SessionsApi* | [**SessionDelete**](docs/SessionsApi.md#sessiondelete) | **Delete** /v1/session/{session_id} | Delete a Session
*SessionsApi* | [**SessionPost**](docs/SessionsApi.md#sessionpost) | **Post** /v1/session | Create a Session
*SessionsApi* | [**SessionsBulkStop**](docs/SessionsApi.md#sessionsbulkstop) | **Post** /v1/sessions/bulk-stop | Delete Sessions in Bulk
*TelemetryApi* | [**ActiveDeploymentTelemetryGet**](docs/TelemetryApi.md#activedeploymenttelemetryget) | **Get** /v1/telemetry/active-deployments/{retrieval_key} | Get the Result of an Active Deployment Telemetry Request
*TelemetryApi* | [**ActiveDeploymentTelemetryPost**](docs/TelemetryApi.md#activedeploymenttelemetrypost) | **Post** /v1/telemetry/active-deployments | Create a New Active Deployment Telemetry Request


## Documentation For Models

 - [ActiveDeploymentTelemetryGetResult](docs/ActiveDeploymentTelemetryGetResult.md)
 - [ActiveDeploymentTelemetryRequest](docs/ActiveDeploymentTelemetryRequest.md)
 - [ActiveDeploymentTelemetryResponse](docs/ActiveDeploymentTelemetryResponse.md)
 - [ActiveDeploymentTelemetryScore](docs/ActiveDeploymentTelemetryScore.md)
 - [ApiModelContainercrashdata](docs/ApiModelContainercrashdata.md)
 - [ApiModelContainerlogs](docs/ApiModelContainerlogs.md)
 - [ApiModelDeploymentfilter](docs/ApiModelDeploymentfilter.md)
 - [ApiModelLocation](docs/ApiModelLocation.md)
 - [ApiModelLocationbeacon](docs/ApiModelLocationbeacon.md)
 - [AppVersionCreateResponse](docs/AppVersionCreateResponse.md)
 - [AppVersionCreateSessionConfig](docs/AppVersionCreateSessionConfig.md)
 - [AppVersionDelete](docs/AppVersionDelete.md)
 - [AppVersionEnv](docs/AppVersionEnv.md)
 - [AppVersionList](docs/AppVersionList.md)
 - [AppVersionPayload](docs/AppVersionPayload.md)
 - [AppVersionPort](docs/AppVersionPort.md)
 - [AppVersionProbe](docs/AppVersionProbe.md)
 - [AppVersionUpdatePayload](docs/AppVersionUpdatePayload.md)
 - [AppVersionUpdateResponse](docs/AppVersionUpdateResponse.md)
 - [AppVersionUpdateSessionConfig](docs/AppVersionUpdateSessionConfig.md)
 - [AppVersionWhitelistEntry](docs/AppVersionWhitelistEntry.md)
 - [AppVersionWhitelistEntryPayload](docs/AppVersionWhitelistEntryPayload.md)
 - [AppVersionWhitelistEntrySuccess](docs/AppVersionWhitelistEntrySuccess.md)
 - [AppVersionWhitelistResponse](docs/AppVersionWhitelistResponse.md)
 - [Application](docs/Application.md)
 - [ApplicationPatch](docs/ApplicationPatch.md)
 - [ApplicationPost](docs/ApplicationPost.md)
 - [Applications](docs/Applications.md)
 - [BaseModel](docs/BaseModel.md)
 - [BulkSessionDelete](docs/BulkSessionDelete.md)
 - [BulkSessionPost](docs/BulkSessionPost.md)
 - [ClientRelayPort](docs/ClientRelayPort.md)
 - [ComponentCredentials](docs/ComponentCredentials.md)
 - [ContainerLogStorageModel](docs/ContainerLogStorageModel.md)
 - [ContextDeploymentTagResponse](docs/ContextDeploymentTagResponse.md)
 - [CustomBulkSessionModel](docs/CustomBulkSessionModel.md)
 - [CustomBulkSessionsModel](docs/CustomBulkSessionsModel.md)
 - [CustomSessionDeleteModel](docs/CustomSessionDeleteModel.md)
 - [CustomSessionModel](docs/CustomSessionModel.md)
 - [Delete](docs/Delete.md)
 - [DeleteRequestReceived](docs/DeleteRequestReceived.md)
 - [DeployEnvModel](docs/DeployEnvModel.md)
 - [DeployModel](docs/DeployModel.md)
 - [Deployment](docs/Deployment.md)
 - [DeploymentBulkStopFiltersPayload](docs/DeploymentBulkStopFiltersPayload.md)
 - [DeploymentBulkStopPayload](docs/DeploymentBulkStopPayload.md)
 - [DeploymentBulkStopResponse](docs/DeploymentBulkStopResponse.md)
 - [DeploymentListData](docs/DeploymentListData.md)
 - [DeploymentLocation](docs/DeploymentLocation.md)
 - [DeploymentSessionContext](docs/DeploymentSessionContext.md)
 - [DeploymentStopResponse](docs/DeploymentStopResponse.md)
 - [DeploymentTagListResponse](docs/DeploymentTagListResponse.md)
 - [DeploymentTagPayload](docs/DeploymentTagPayload.md)
 - [DeploymentTagResponse](docs/DeploymentTagResponse.md)
 - [DeploymentUpdatePayload](docs/DeploymentUpdatePayload.md)
 - [DeploymentUpdateResponse](docs/DeploymentUpdateResponse.md)
 - [Deployments](docs/Deployments.md)
 - [EndpointStorageDeleteResponse](docs/EndpointStorageDeleteResponse.md)
 - [EndpointStorageGetResponse](docs/EndpointStorageGetResponse.md)
 - [EndpointStorageListResponse](docs/EndpointStorageListResponse.md)
 - [EndpointStoragePatchPayload](docs/EndpointStoragePatchPayload.md)
 - [EndpointStoragePatchResponse](docs/EndpointStoragePatchResponse.md)
 - [EndpointStoragePostPayload](docs/EndpointStoragePostPayload.md)
 - [EndpointStoragePostResponse](docs/EndpointStoragePostResponse.md)
 - [FleetDeleteResponse](docs/FleetDeleteResponse.md)
 - [FleetGetResponse](docs/FleetGetResponse.md)
 - [FleetList](docs/FleetList.md)
 - [FleetPatchPayload](docs/FleetPatchPayload.md)
 - [FleetPatchResponse](docs/FleetPatchResponse.md)
 - [FleetPoliciesGetResponse](docs/FleetPoliciesGetResponse.md)
 - [FleetPoliciesPatchPayload](docs/FleetPoliciesPatchPayload.md)
 - [FleetPoliciesPostPayload](docs/FleetPoliciesPostPayload.md)
 - [FleetPoliciesPostResponse](docs/FleetPoliciesPostResponse.md)
 - [FleetPostPayload](docs/FleetPostPayload.md)
 - [FleetPostResponse](docs/FleetPostResponse.md)
 - [GeoIpListModel](docs/GeoIpListModel.md)
 - [HorizontalScalerAppVersionLink](docs/HorizontalScalerAppVersionLink.md)
 - [HorizontalScalerConstraintList](docs/HorizontalScalerConstraintList.md)
 - [IpAddressLookupLocation](docs/IpAddressLookupLocation.md)
 - [IpAddressLookupLocationContinent](docs/IpAddressLookupLocationContinent.md)
 - [IpAddressLookupLocationCountry](docs/IpAddressLookupLocationCountry.md)
 - [IpAddressLookupResponse](docs/IpAddressLookupResponse.md)
 - [IpAddressResponse](docs/IpAddressResponse.md)
 - [IpAddressesLookupPayload](docs/IpAddressesLookupPayload.md)
 - [IpAddressesLookupResponse](docs/IpAddressesLookupResponse.md)
 - [LobbyCreatePayload](docs/LobbyCreatePayload.md)
 - [LobbyDeployPayload](docs/LobbyDeployPayload.md)
 - [LobbyReadResponse](docs/LobbyReadResponse.md)
 - [LobbyTerminatePayload](docs/LobbyTerminatePayload.md)
 - [Location](docs/Location.md)
 - [LocationBeaconList](docs/LocationBeaconList.md)
 - [LocationModel](docs/LocationModel.md)
 - [Locations](docs/Locations.md)
 - [MappedPortResponse](docs/MappedPortResponse.md)
 - [MatchmakerComponentCreate](docs/MatchmakerComponentCreate.md)
 - [MatchmakerComponentEnvListResponse](docs/MatchmakerComponentEnvListResponse.md)
 - [MatchmakerComponentEnvsCreate](docs/MatchmakerComponentEnvsCreate.md)
 - [MatchmakerComponentEnvsResponse](docs/MatchmakerComponentEnvsResponse.md)
 - [MatchmakerComponentEnvsUpdate](docs/MatchmakerComponentEnvsUpdate.md)
 - [MatchmakerComponentListResponse](docs/MatchmakerComponentListResponse.md)
 - [MatchmakerComponentResponse](docs/MatchmakerComponentResponse.md)
 - [MatchmakerComponentResponseCredentials](docs/MatchmakerComponentResponseCredentials.md)
 - [MatchmakerComponentUpdate](docs/MatchmakerComponentUpdate.md)
 - [MatchmakerCreate](docs/MatchmakerCreate.md)
 - [MatchmakerListResponse](docs/MatchmakerListResponse.md)
 - [MatchmakerManagedReleaseCreate](docs/MatchmakerManagedReleaseCreate.md)
 - [MatchmakerManagedReleaseResponse](docs/MatchmakerManagedReleaseResponse.md)
 - [MatchmakerManagedReleaseUpdate](docs/MatchmakerManagedReleaseUpdate.md)
 - [MatchmakerReleaseConfigCreate](docs/MatchmakerReleaseConfigCreate.md)
 - [MatchmakerReleaseConfigResponse](docs/MatchmakerReleaseConfigResponse.md)
 - [MatchmakerReleaseConfigUpdate](docs/MatchmakerReleaseConfigUpdate.md)
 - [MatchmakerReleaseCreate](docs/MatchmakerReleaseCreate.md)
 - [MatchmakerReleaseCreateBase](docs/MatchmakerReleaseCreateBase.md)
 - [MatchmakerReleaseResponse](docs/MatchmakerReleaseResponse.md)
 - [MatchmakerReleaseResponseBase](docs/MatchmakerReleaseResponseBase.md)
 - [MatchmakerReleaseUpdate](docs/MatchmakerReleaseUpdate.md)
 - [MatchmakerReleaseUpdateBase](docs/MatchmakerReleaseUpdateBase.md)
 - [MatchmakerResponse](docs/MatchmakerResponse.md)
 - [MatchmakerUpdate](docs/MatchmakerUpdate.md)
 - [MetricsModel](docs/MetricsModel.md)
 - [MetricsResponse](docs/MetricsResponse.md)
 - [ModelError](docs/ModelError.md)
 - [Monitor](docs/Monitor.md)
 - [NetworkMetricsModel](docs/NetworkMetricsModel.md)
 - [Pagination](docs/Pagination.md)
 - [Paginator](docs/Paginator.md)
 - [PatchSessionModel](docs/PatchSessionModel.md)
 - [Payload](docs/Payload.md)
 - [PortMapping](docs/PortMapping.md)
 - [PullProfileAppVersionLinkResponse](docs/PullProfileAppVersionLinkResponse.md)
 - [PullProfileGetResponse](docs/PullProfileGetResponse.md)
 - [PullProfilePatchPayload](docs/PullProfilePatchPayload.md)
 - [PullProfilePostPayload](docs/PullProfilePostPayload.md)
 - [PullProfilePostResponse](docs/PullProfilePostResponse.md)
 - [PullProfilesListResponse](docs/PullProfilesListResponse.md)
 - [PulloProfilePatchResponse](docs/PulloProfilePatchResponse.md)
 - [RelayFilterModel](docs/RelayFilterModel.md)
 - [RelayResponse](docs/RelayResponse.md)
 - [RelaySessionBaseResponse](docs/RelaySessionBaseResponse.md)
 - [RelaySessionBaseResponseRelay](docs/RelaySessionBaseResponseRelay.md)
 - [RelaySessionCreatePayload](docs/RelaySessionCreatePayload.md)
 - [RelaySessionListResponse](docs/RelaySessionListResponse.md)
 - [RelaySessionUser](docs/RelaySessionUser.md)
 - [RelaySessionUserBaseResponse](docs/RelaySessionUserBaseResponse.md)
 - [RelaySessionUserBaseResponseSessionUser](docs/RelaySessionUserBaseResponseSessionUser.md)
 - [RelaySessionUserResponse](docs/RelaySessionUserResponse.md)
 - [RelayUserAuthorizePayload](docs/RelayUserAuthorizePayload.md)
 - [RelayUserRevokePayload](docs/RelayUserRevokePayload.md)
 - [Request](docs/Request.md)
 - [SelectorEnvModel](docs/SelectorEnvModel.md)
 - [SelectorModel](docs/SelectorModel.md)
 - [SelectorModelEnv](docs/SelectorModelEnv.md)
 - [ServerRelayPort](docs/ServerRelayPort.md)
 - [SessionBulkStopFiltersPayload](docs/SessionBulkStopFiltersPayload.md)
 - [SessionBulkStopPayload](docs/SessionBulkStopPayload.md)
 - [SessionBulkStopResponse](docs/SessionBulkStopResponse.md)
 - [SessionContext](docs/SessionContext.md)
 - [SessionDelete](docs/SessionDelete.md)
 - [SessionFilterModel](docs/SessionFilterModel.md)
 - [SessionGet](docs/SessionGet.md)
 - [SessionModel](docs/SessionModel.md)
 - [SessionRequest](docs/SessionRequest.md)
 - [SessionStopResponse](docs/SessionStopResponse.md)
 - [SessionUser](docs/SessionUser.md)
 - [SessionUserContext](docs/SessionUserContext.md)
 - [Sessions](docs/Sessions.md)
 - [Status](docs/Status.md)
 - [Tag](docs/Tag.md)
 - [TotalMetricsModel](docs/TotalMetricsModel.md)


## Documentation For Authorization

## apiKey
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

info@edgegap.com

