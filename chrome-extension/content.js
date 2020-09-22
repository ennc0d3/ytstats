// content.js
// This is the javascript that can run in the context of
// the webpages and aswell interact with the pages.
//
//

chrome.runtime.onMessage.addListener(
  function (request, sender, sendResponse) {
    if (request.message === "clicked_browser_action") {
      console.log("Youtube stats on tab!");
      var firstHref = "https://ennc0d3.github.io"
      $.urlParam = function(name,url){
        var results = new RegExp('[\?&]' + name + '=([^&#]*)').exec(url);
        console.log("parms:" + results)
       //return results[0] || 0;
        return "fixfunc"
      }
      var videoID = $.urlParam('v', request.url)
      console.log("Search for the video id:" + videoID)
      videoID = "fixme"
      var firstHref = "https://ennc0d3.github.io"
      chrome.runtime.sendMessage({ "message": "open_new_tab",
            "url": firstHref ,
            "vid": videoID });
    }
  }
);
