---
description: Custom Axle Component
---

# Right Side Bar

![Right Side Bar Finished State](../.gitbook/assets/right-side-bar-finished.png)

{% tabs %}
{% tab title="Usage" %}
### Usage

The Right Side Bar shows summary information about a transaction.
{% endtab %}

{% tab title="Requirements" %}
### Requirements

* The Right Side Bar must be updated whenever values are saved or calculations are performed as a part of the transaction.

### Data Mappings

#### Personal Information

| Field | Source |
| :--- | :--- |
| Email | Member Input |
| Phone | Member Input |
| Used/Line | Borrower Credit Line |
| Interest Rate | Borrower Credit Line |

#### Asset Name

| Field | Source |
| :--- | :--- |
| Year Make Model | Carputty Database |
| Mileage | Member Input, Odometer Verification Drawer |
| Current Valuation | Calculated by Monroney Drawer and Odometer Verification Drawer |
| Remaining Payment | Member Input |

#### Transaction

| Field | Source |
| :--- | :--- |
| Residual Amount | Member Input, Final Statement Form |
| Maximum Allowed | Calculated |
| Advance Amount | Calculated |
| Transaction Total | Calculated |
| Minimum Down Payment | Calculated |

#### Valuation

| Field | Source |
| :--- | :--- |
| Preliminary Valuation | Monroney Label Drawer |
| Final Valuation | Odometer Verification Drawer |

### Fee Total

| Field | Source |
| :--- | :--- |
| Estimated | Preliminary Statement Form |
| Final | Final Statement Form |

### Lessor

| Field | Source |
| :--- | :--- |
| Lease Account \# | Member Input |
| Type | Preliminary Statement Form |
| Phone | Preliminary Statement Form |
| Address | Preliminary Statement Form |
{% endtab %}

{% tab title="Interactions" %}
* Jonathan Customer name is a link to the member's profile page
* Personal information is a link to the member's profile page

![](../.gitbook/assets/interaction.png)
{% endtab %}

{% tab title="States" %}
### States

### On-Load

![](../.gitbook/assets/right-side-bar-onload.png)

### Trim

![](../.gitbook/assets/right-side-bar-trim.png)

### Preliminary Statement

![](../.gitbook/assets/right-side-bar-prelim.png)

### Final Statement

![](../.gitbook/assets/right-side-bar-final.png)
{% endtab %}

{% tab title="Code" %}

{% endtab %}
{% endtabs %}









