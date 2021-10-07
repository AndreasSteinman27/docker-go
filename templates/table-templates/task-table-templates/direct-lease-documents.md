---
description: Variant of Task Table
---

# Direct Lease Documents

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=624%3A23893" %}



{% tabs %}
{% tab title="Overview" %}
### Usage

The Direct Lease Documents table tracks the documents and verification for a Direct lease buyout transaction.

### Base Component

{% page-ref page="../../../components/task-tables/task-table.md" %}
{% endtab %}

{% tab title="Functional Requirements" %}
#### Fields

Fields marked as Required must be completed and validated to submit the form attached to this table.

| Name | Drawer Type | Required |
| :--- | :--- | :--- |
| Odometer Image | Odometer | **Yes** |
| Driver's License - Front | View/Upload | **Yes** |
| Driver's License - Back | View/Upload | **Yes** |
| Proof of Insurance | View/Upload | **Yes** |
| Notarized Power of Attorney | View/Upload | **NO** - but a red asterisk should still appear.  |

### Interactions

See [Task Table](../../../components/task-tables/task-table.md).
{% endtab %}

{% tab title="Validation/Error Handling" %}
### See [Task Table](../../../components/task-tables/task-table.md)
{% endtab %}

{% tab title="Designs" %}
### On-Loading

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=624%3A23893" %}

### Status Change when tasks are complete

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=1971%3A15710" %}

### Complete

{% embed url="https://www.figma.com/file/w78ZiMR2USgl1CwXVrcxXv/?node-id=876%3A26057" %}
{% endtab %}

{% tab title="Code" %}

{% endtab %}
{% endtabs %}



