# Upload/Verify Drawer

![](../../.gitbook/assets/screen-shot-2021-09-29-at-4.27.10-pm.png)

{% tabs %}
{% tab title="Overview" %}
This component is used by MSAs to upload documents and verify existing documents.

### Components with Interactions:

{% page-ref page="../../ingredients/task-tables/task-table.md" %}



### Tasks that use this Drawer Component <a id="tasks-that-use-this-drawer-component"></a>

| Name | Task Table |
| :--- | :--- |
| Driver's License Front | ​[Direct Lease Documents](https://app.gitbook.com/@carputty/s/axle/~/drafts/-MknK5R5yIP9tXv-bbTE/components/task-tables/task-table/direct-lease-documents)​ |
| Driver's License Back | ​[Direct Lease Documents](https://app.gitbook.com/@carputty/s/axle/~/drafts/-MknK5R5yIP9tXv-bbTE/components/task-tables/task-table/direct-lease-documents)​ |
| Proof of Insurance | ​[Direct Lease Documents](https://app.gitbook.com/@carputty/s/axle/~/drafts/-MknK5R5yIP9tXv-bbTE/components/task-tables/task-table/direct-lease-documents)​ |
| Notarized Attorney | ​[Direct Lease Documents](https://app.gitbook.com/@carputty/s/axle/~/drafts/-MknK5R5yIP9tXv-bbTE/components/task-tables/task-table/direct-lease-documents)​ |
| Bill of Sale | ​[Disbursement](https://app.gitbook.com/@carputty/s/axle/~/drafts/-MknK5R5yIP9tXv-bbTE/components/task-tables/task-table/disbursement)​ |
| Title Application | ​[Disbursement](https://app.gitbook.com/@carputty/s/axle/~/drafts/-MknK5R5yIP9tXv-bbTE/components/task-tables/task-table/disbursement)​ |
| Insurance Coverage Verification | ​[Disbursement](https://app.gitbook.com/@carputty/s/axle/~/drafts/-MknK5R5yIP9tXv-bbTE/components/task-tables/task-table/disbursement)​ |
| Insurance Loss Payee Verification | ​[Disbursement](https://app.gitbook.com/@carputty/s/axle/~/drafts/-MknK5R5yIP9tXv-bbTE/components/task-tables/task-table/disbursement)​ |
| Engagement Letter | ​[Disbursement](https://app.gitbook.com/@carputty/s/axle/~/drafts/-MknK5R5yIP9tXv-bbTE/components/task-tables/task-table/disbursement)​ |
| Odometer Verification |  |
{% endtab %}

{% tab title="Functional Requirements" %}
### **Requirements:**

* The MSA should be able to upload an image.
  * Files that are not PDF, JPG, TIFF, PNG, or GIF are not supported.
  * Files over 10MB aren’t supported
* If an image has already been uploaded, and a new image is uploaded to replace it, the old image should be archived/hidden, and the new image will replace the old given the type of document selected.
* Upon an upload being initiated, the file should be validated.
* The upload drawer can only accept one file at a time
* The Save button should be inactive by default.
* The MSA should be able to view the uploaded image
* The MSA must be able to mark the image as pass or fail for verification purposes. The Save button should only be activated once pass or fail has been selected.
* If the user clicks save, the uploaded file should be stored as the document type that it was uploaded for. If an image was already submitted for this document type, it should be archived and hidden. The Verification status should be updated to match the verification status given.
* If discard or close is clicked, any uploads or verification changes should be discarded.

### **Steps:**

1. The drawer is opened
2. A file is uploaded to the drawer
3. The file is displayed in the drawer
4. The verification status of the document is marked \(pass/fail\)
5. The upload and verification status are saved.

### UI Interactions

**Upload Component**

* If a valid file has already been uploaded, it should be displayed.
* The upload component should accept uploads either from the system file picker or from drag/drop.
  * If the user clicks the upload component, the system file picker should appear. Once a file is selected, an upload should be initiated with the selected file.
  * If the user drags a file to the upload component, an upload should be initiated with the selected file.
* Upon an upload being initiated, the file should be validated.
  * File must be PDF, JPG, TIFF, PNG, or GIF and under 10MB.
* Once a successful upload is performed, the image should be displayed in the drawer, replacing any existing images if present.
* If the user uploads a file when another file was already uploaded, the old upload should be discarded and the new upload should replace the existing upload upon save.
  * The upload drawer can only accept one file at a time.

**Visual Inspection**

* If the Visual Inspection dropdown is null, the Save button should be inactive.
* If the MSA selects Pass or Fail on the Visual Inspection dropdown, the Save button should be activated.

**Save Button**

* The Save button's default state is inactive.
  * If it is clicked while inactive, nothing should happen.
* If the Save Button is clicked while active, the associated document should be updated with the appropriate status, the previously uploaded image \(if applicable\) should be hidden/archived, the new document should replace the old document for the selected document type, and the drawer should close.
  * Verified if the Visual Inspection is Pass
  * Failed if the Visual Inspection is Fail

**Discard and Close Buttons**

* The Discard button should always be active. When clicked, it should discard any uploads or input received, and the drawer should close.
  * The close button behaves identically.
{% endtab %}

{% tab title="Validation/Error Handling" %}
Error message \(onUpload\)

* IF image is not valid image type then error is shown:
  * Invalid File Type
* IF image is too large then error is shown:
  * File is over 10MB

If file validation passes, the upload should begin. A progress bar should be shown with the upload.
{% endtab %}

{% tab title="Designs" %}
#### Task Table

![](../../.gitbook/assets/screen-shot-2021-09-29-at-4.11.00-pm.png)

#### Drawer OnLoad - Empty

![](../../.gitbook/assets/screen-shot-2021-09-29-at-4.23.12-pm.png)

#### 

#### Drawer OnLoad - Upload Exists

![](../../.gitbook/assets/screen-shot-2021-09-29-at-4.27.10-pm.png)

#### ERROR STATE IMAGE TBD

#### Complete

![](../../.gitbook/assets/screen-shot-2021-09-29-at-4.30.50-pm.png)
{% endtab %}

{% tab title="Code" %}

{% endtab %}
{% endtabs %}

### 

