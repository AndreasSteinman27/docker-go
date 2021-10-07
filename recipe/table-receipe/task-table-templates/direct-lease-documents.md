# Direct Lease Documents

![](../../../.gitbook/assets/direct-lease-documents-onload%20%281%29.png)

{% tabs %}
{% tab title="Overview" %}
### Usage

The Direct Lease Documents table tracks the documents and verification for a Direct lease buyout transaction.

### Base Component

{% page-ref page="../../../ingredients/task-tables/task-table.md" %}
{% endtab %}

{% tab title="Functional Requirements" %}
#### Fields

Fields marked as Required must be completed and validated to submit the form attached to this table.

| Name | Drawer Type | Required |
| :--- | :--- | :--- |
| Odometer Image | Odometer | **Yes** |
| Driver's License - Front | Upload/Verify | **Yes** |
| Driver's License - Back | Upload/Verify | **Yes** |
| Proof of Insurance | Upload/Verify | **Yes** |
| Notarized Power of Attorney | Upload/Verify | **NO** - but a red asterisk should still appear.  |

### Interactions

See [Task Table](../../../ingredients/task-tables/task-table.md).
{% endtab %}

{% tab title="Validation/Error Handling" %}
### See [Task Table](../../../ingredients/task-tables/task-table.md)
{% endtab %}

{% tab title="States" %}
### On-Loading

![](../../../.gitbook/assets/direct-lease-documents-onload%20%282%29.png)

### Status Change when tasks are complete

![](../../../.gitbook/assets/direct-lease-documents-status-change.png)

### Complete

![](../../../.gitbook/assets/direct-lease-documents-complete%20%281%29.png)
{% endtab %}

{% tab title="Code" %}

{% endtab %}
{% endtabs %}





