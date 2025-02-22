---
description: Custom Axle Component
---

# Right Side Bar

![Right Side Bar Finished State](../.gitbook/assets/right-side-bar-finished.png)

{% tabs %}
{% tab title="Overview" %}
### Usage

The Right Side Bar shows summary information about a transaction.

### Components with Interactions:

| Page |
| :--- |
| [Monroney Label Drawer](../recipe/drawer/drawer-trim.md) |
| [Preliminary Statement](../recipe/form-receipe/statement-input-form-templates/lbo-preliminary-statement.md) |
| [Final Statement](../recipe/form-receipe/statement-input-form-templates/lbo-final-statement.md) |
{% endtab %}

{% tab title="Functional Requirements" %}
### Requirements

* The Right Side Bar must be updated whenever values are saved or calculations are performed as a part of the transaction.
  * This happens when the Prelim form is saved, when the Final form is saved, when the Trim is verified via the Monroney Label Drawer, and when the Odometer is verified in the Odometer drawer.

### Interactions

* Clicking the name of the member should bring the user to the member's Member page.
* Clicking Personal Information should bring the user to the member's Member page.

### Steps:

1. Verify the Trim of a vehicle, and note that the prelim valuation is updated.
2. Fill out and save the Preliminary Statement form, and note that the values entered \(and their calculations\) are displayed correctly in the side bar.
3. Verify the Odometer Image of a vehicle, and note that the final valuation is updated.
4. Fill out and save the Final Statement form, and note that the values entered \(and their calculations\) are displayed correctly in the side bar.

### Data Mappings

#### Personal Information

| Field | Format | Source |
| :--- | :--- | :--- |
| Email | Text | Member Input |
| Phone | Phone Number | Member Input |
| Used/Line | Currency / Currency | Borrower Credit Line |
| Interest Rate | Percentage | Borrower Credit Line |

#### Asset Name

| Field | Format | Source |
| :--- | :--- | :--- |
| Year Make Model | Text | Carputty Database |
| Mileage | Mileage | Member Input, Odometer Verification Drawer |
| Current Valuation | Currency | Calculated by Monroney Drawer and Odometer Verification Drawer |
| Remaining Payment | Currency | Member Input |

#### Transaction

| Field | Format | Source |
| :--- | :--- | :--- |
| Residual Amount | Currency | Member Input, Final Statement Form |
| [Maximum Allowed](../overview/calculations.md) | Currency | Calculated |
| [Advance Amount](../overview/calculations.md) | Currency | Calculated |
| [Transaction Total](../overview/calculations.md) | Currency | Calculated |
| [Minimum Down Payment](../overview/calculations.md) | Currency | Calculated |

#### Valuation

| Field | Format | Source |
| :--- | :--- | :--- |
| Preliminary Valuation | Currency | Monroney Label Drawer |
| Final Valuation | Currency | Odometer Verification Drawer |

### Fee Total

| Field | Format | Source |
| :--- | :--- | :--- |
| Estimated | Currency | Preliminary Statement Form |
| Final | Currency | Final Statement Form |

### Lessor

| Field | Format | Source |
| :--- | :--- | :--- |
| Lease Account \# | Text | Member Input |
| Type | Format | Preliminary Statement Form |
| Phone | Phone Number | Preliminary Statement Form |
| Address | Text | Preliminary Statement Form |
{% endtab %}

{% tab title="Validation/Error Handling" %}
* This sidebar does not show errors, however, the calculations should match the values we have documented, this must be verified in QA.
{% endtab %}

{% tab title="Designs" %}
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









