# Upload

![](../../.gitbook/assets/upload_defaultstate.png)

{% tabs %}
{% tab title="Usage" %}
The Upload component is used to allow the user to upload files.

### Base Component

Ant Design's Upload class with Drag and Drop functionality: [https://ant.design/components/upload/\#components-upload-demo-drag](https://ant.design/components/upload/#components-upload-demo-drag)
{% endtab %}

{% tab title="Requirements" %}
### Requirements

* If a user drags and drops a file to the Upload component, the file should be uploaded. 
* If the Upload Component is clicked, the system file picker should appear.
* If an upload is in progress, a progress bar should show the state of the upload. 
* The component must return an error if the file does not meet size or file type requirements \(for example, the [Upload/Verify Drawer](../../recipe/drawer/upload-verify-drawer.md) requires JPG/PDF/PNG/GIF/TIFF and &lt;10MB\).
* The component must display an error if the upload fails.
{% endtab %}

{% tab title="Error Handling" %}
* Invalid File Type
* The file is over XMB
  * X = maximum file size
{% endtab %}
{% endtabs %}

