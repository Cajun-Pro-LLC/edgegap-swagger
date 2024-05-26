/*
 * Edgegap V1
 *
 * <h1 style=\"margin-top: 2rem\">Introduction</h1> <p>The Edgegap API lets you manage all your resources from standard HTTP requests. We promote using them to automate all your processes that are using Edgegap.</p> <p>If you have any questions, don't hesitate to contact us via email, or you can also jump on our <a href=\"https://discord.com/invite/GYaHcKR9a5\" target=\"_blank\">Discord</a>. We will be happy to help. Feel free to make features request; we also love those.</p> <div class=\"theme-doc-markdown markdown\"><h2 class=\"anchor anchorWithStickyNavbar_node_modules-@docusaurus-theme-classic-lib-theme-Heading-styles-module\" id=\"pagination---response\">Pagination - Response<a class=\"hash-link\" href=\"#pagination---response\" title=\"Direct link to heading\">​</a></h2><p>The GET response body can be returned with pagination to avoid requesting all the data at once.</p><p>Pagination object will always be under the  <em>paginator</em> key.</p><div class=\"language-json codeBlockContainer_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Container-styles-module theme-code-block\" style=\"--prism-color:#bfc7d5; --prism-background-color:#292d3e;\"><div class=\"codeBlockContent_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><pre tabindex=\"0\" class=\"prism-code language-json codeBlock_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module thin-scrollbar\"><code class=\"codeBlockLines_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">{</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"number\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"The Current page, default=1\"</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"next_page_number\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"The Next page number or null\"</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"previous_page_number\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"The Previous page number or null\"</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"paginator\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">{</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">    </span><span class=\"token property\">\"num_pages\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"The total numbers of pages\"</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">}</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"has_next\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"Boolean if there is a next page\"</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"has_previous\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"Boolean if there is a previous page\"</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\"></span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">}</span><br></span></code></pre><div class=\"buttonGroup_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><button type=\"button\" aria-label=\"Copy code to clipboard\" title=\"Copy\" class=\"clean-btn\"><span class=\"copyButtonIcons_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" aria-hidden=\"true\"><svg class=\"copyButtonIcon_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" viewBox=\"0 0 24 24\"><path d=\"M19,21H8V7H19M19,5H8A2,2 0 0,0 6,7V21A2,2 0 0,0 8,23H19A2,2 0 0,0 21,21V7A2,2 0 0,0 19,5M16,1H4A2,2 0 0,0 2,3V17H4V3H16V1Z\"></path></svg><svg class=\"copyButtonSuccessIcon_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" viewBox=\"0 0 24 24\"><path d=\"M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z\"></path></svg></span></button></div></div></div><p>Full Body Example:</p><div class=\"language-json codeBlockContainer_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Container-styles-module theme-code-block\" style=\"--prism-color:#bfc7d5; --prism-background-color:#292d3e;\"><div class=\"codeBlockContent_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><pre tabindex=\"0\" class=\"prism-code language-json codeBlock_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module thin-scrollbar\"><code class=\"codeBlockLines_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">{</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"count\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token number\" style=\"color: rgb(247, 140, 108);\">100</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"data\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">[</span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"value-0\"</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"> </span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"[...]\"</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"> </span><span class=\"token string\" style=\"color: rgb(195, 232, 141);\">\"value-9\"</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">]</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"success\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token boolean\" style=\"color: rgb(255, 88, 116);\">true</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token property\">\"pagination\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">{</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">    </span><span class=\"token property\">\"number\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token number\" style=\"color: rgb(247, 140, 108);\">1</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">    </span><span class=\"token property\">\"next_page_number\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token number\" style=\"color: rgb(247, 140, 108);\">2</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">    </span><span class=\"token property\">\"previous_page_number\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token null keyword\" style=\"font-style: italic;\">null</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">    </span><span class=\"token property\">\"paginator\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">{</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">      </span><span class=\"token property\">\"num_pages\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token number\" style=\"color: rgb(247, 140, 108);\">10</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">    </span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">}</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">    </span><span class=\"token property\">\"has_next\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token boolean\" style=\"color: rgb(255, 88, 116);\">true</span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">,</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">    </span><span class=\"token property\">\"has_previous\"</span><span class=\"token operator\" style=\"color: rgb(137, 221, 255);\">:</span><span class=\"token plain\"> </span><span class=\"token boolean\" style=\"color: rgb(255, 88, 116);\">false</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">  </span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">}</span><span class=\"token plain\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\"></span><span class=\"token punctuation\" style=\"color: rgb(199, 146, 234);\">}</span><br></span></code></pre><div class=\"buttonGroup_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><button type=\"button\" aria-label=\"Copy code to clipboard\" title=\"Copy\" class=\"clean-btn\"><span class=\"copyButtonIcons_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" aria-hidden=\"true\"><svg class=\"copyButtonIcon_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" viewBox=\"0 0 24 24\"><path d=\"M19,21H8V7H19M19,5H8A2,2 0 0,0 6,7V21A2,2 0 0,0 8,23H19A2,2 0 0,0 21,21V7A2,2 0 0,0 19,5M16,1H4A2,2 0 0,0 2,3V17H4V3H16V1Z\"></path></svg><svg class=\"copyButtonSuccessIcon_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" viewBox=\"0 0 24 24\"><path d=\"M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z\"></path></svg></span></button></div></div></div><h2 class=\"anchor anchorWithStickyNavbar_node_modules-@docusaurus-theme-classic-lib-theme-Heading-styles-module\" id=\"pagination---parameters\">Pagination - Parameters<a class=\"hash-link\" href=\"#pagination---parameters\" title=\"Direct link to heading\">​</a></h2><p>You can add those values to manipulate the pagination object in the URL Parameters.</p><ul class=\"\"><li>page</li><li>limit</li></ul><p>Example:</p><div class=\"language-text codeBlockContainer_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Container-styles-module theme-code-block\" style=\"--prism-color:#bfc7d5; --prism-background-color:#292d3e;\"><div class=\"codeBlockContent_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><pre tabindex=\"0\" class=\"prism-code language-text codeBlock_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module thin-scrollbar\"><code class=\"codeBlockLines_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\"># To get the second page</span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">GET - https://api.edgegap.com/v1/apps?page=2</span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\" style=\"display: inline-block;\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\"># To change the count of element in one page (20/page)</span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">GET - https://api.edgegap.com/v1/apps?limit=20</span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\" style=\"display: inline-block;\"></span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\"># You can mix those (20/page, second page)</span><br></span><span class=\"token-line\" style=\"color: rgb(191, 199, 213);\"><span class=\"token plain\">GET - https://api.edgegap.com/v1/apps?page=2&amp;limit=20</span><br></span></code></pre><div class=\"buttonGroup_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-Content-styles-module\"><button type=\"button\" aria-label=\"Copy code to clipboard\" title=\"Copy\" class=\"clean-btn\"><span class=\"copyButtonIcons_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" aria-hidden=\"true\"><svg class=\"copyButtonIcon_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" viewBox=\"0 0 24 24\"><path d=\"M19,21H8V7H19M19,5H8A2,2 0 0,0 6,7V21A2,2 0 0,0 8,23H19A2,2 0 0,0 21,21V7A2,2 0 0,0 19,5M16,1H4A2,2 0 0,0 2,3V17H4V3H16V1Z\"></path></svg><svg class=\"copyButtonSuccessIcon_node_modules-@docusaurus-theme-classic-lib-theme-CodeBlock-CopyButton-styles-module\" viewBox=\"0 0 24 24\"><path d=\"M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z\"></path></svg></span></button></div></div></div></div>
 *
 * API version: 1.0.0
 * Contact: info@edgegap.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type CustomSessionsApiService service

/*
CustomSessionsApiService Delete a Custom Session
Delete a custom session. Once deleted, a custom session is no more accessible and does not have a history. The deployment associated will not be deleted.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param customId Custom ID Managed by you
  - @param requestId Deployment Request ID

@return SessionDelete
*/
func (a *CustomSessionsApiService) DeleteCustomSession(ctx context.Context, customId string, requestId string) (SessionDelete, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Delete")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SessionDelete
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/deployment/{request_id}/custom/session/{custom_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom_id"+"}", fmt.Sprintf("%v", customId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"request_id"+"}", fmt.Sprintf("%v", requestId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v SessionDelete
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
CustomSessionsApiService Delete Custom Sessions in Bulk
Delete multiple custom sessions from a specific deployment.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param requestId Deployment Request ID
  - @param payload

@return BulkSessionDelete
*/
func (a *CustomSessionsApiService) DeleteCustomSessions(ctx context.Context, requestId string, payload CustomSessionDeleteModel) (BulkSessionDelete, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Delete")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue BulkSessionDelete
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/deployment/{request_id}/custom/sessions"
	localVarPath = strings.Replace(localVarPath, "{"+"request_id"+"}", fmt.Sprintf("%v", requestId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &payload
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v BulkSessionDelete
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
CustomSessionsApiService Get a Custom Session
Retrieve the information of a custom session.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param customId Custom ID Managed by you
  - @param requestId Deployment Request ID

@return SessionGet
*/
func (a *CustomSessionsApiService) GetCustomSession(ctx context.Context, customId string, requestId string) (SessionGet, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SessionGet
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/deployment/{request_id}/custom/session/{custom_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom_id"+"}", fmt.Sprintf("%v", customId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"request_id"+"}", fmt.Sprintf("%v", requestId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v SessionGet
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
CustomSessionsApiService List All Custom Session of a Deployment
List all custom sessions of a deployment.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param requestId Deployment Request ID

@return Sessions
*/
func (a *CustomSessionsApiService) GetCustomSessions(ctx context.Context, requestId string) (Sessions, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Sessions
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/deployment/{request_id}/custom/sessions"
	localVarPath = strings.Replace(localVarPath, "{"+"request_id"+"}", fmt.Sprintf("%v", requestId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v Sessions
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
CustomSessionsApiService Update a Custom Session
Update a custom session with new specifications.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param customId Custom ID Managed by you
  - @param requestId Deployment Request ID
  - @param payload

@return SessionGet
*/
func (a *CustomSessionsApiService) PatchCustomSession(ctx context.Context, customId string, requestId string, payload CustomSessionModel) (SessionGet, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Patch")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SessionGet
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/deployment/{request_id}/custom/session/{custom_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom_id"+"}", fmt.Sprintf("%v", customId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"request_id"+"}", fmt.Sprintf("%v", requestId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &payload
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v SessionGet
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
CustomSessionsApiService Create a New Custom Session
Create a custom session with users. You must specify a custom ID and a deployment request ID.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param customId Custom ID Managed by you
  - @param requestId Deployment Request ID
  - @param payload

@return SessionRequest
*/
func (a *CustomSessionsApiService) PostCustomSession(ctx context.Context, customId string, requestId string, payload CustomSessionModel) (SessionRequest, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SessionRequest
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/deployment/{request_id}/custom/session/{custom_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom_id"+"}", fmt.Sprintf("%v", customId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"request_id"+"}", fmt.Sprintf("%v", requestId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &payload
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v SessionRequest
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
CustomSessionsApiService Create Custom Sessions in Bulk
Create multiple custom sessions in a deployment. You must specify a custom ID for each.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param requestId Deployment Request ID
  - @param payload

@return BulkSessionPost
*/
func (a *CustomSessionsApiService) PostCustomSessions(ctx context.Context, requestId string, payload CustomBulkSessionsModel) (BulkSessionPost, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue BulkSessionPost
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/deployment/{request_id}/custom/sessions"
	localVarPath = strings.Replace(localVarPath, "{"+"request_id"+"}", fmt.Sprintf("%v", requestId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &payload
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["authorization"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v BulkSessionPost
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
