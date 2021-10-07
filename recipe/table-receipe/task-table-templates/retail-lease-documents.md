# Retail Lease Documents Table

![](../../../.gitbook/assets/retail-lease-documents-onload.png)

{% tabs %}
{% tab title="Usage" %}
### Usage

The Retail Lease Documents table tracks the documents and verification for a Direct lease buyout transaction.

### Base Component

{% page-ref page="../../../ingredients/task-tables/task-table.md" %}
{% endtab %}

{% tab title="Functional Requirements" %}
### Fields

Fields marked as Required must be completed and validated to submit the form attached to this table.

| Name | Drawer Type | Required |
| :--- | :--- | :--- |
| Bill of Sale | Upload/Verify | **Yes** |
| Title Application | Upload/Verify | **Yes** |
| Engagement Letter | Upload/Verify | **Yes** |

### Interactions

See [Task Table](../../../ingredients/task-tables/task-table.md).
{% endtab %}

{% tab title="Validation/Error Handling" %}
### See [Task Table](../../../ingredients/task-tables/task-table.md)
{% endtab %}

{% tab title="States" %}
### States

### On-Loading

![](../../../.gitbook/assets/retail-lease-documents-onload%20%281%29.png)

### Status Changing

![](../../../.gitbook/assets/retail-lease-documents-status-changes.png)

### Complete 

![](../../../.gitbook/assets/retail-lease-documents-complete.png)
{% endtab %}

{% tab title="Code" %}

{% endtab %}
{% endtabs %}







