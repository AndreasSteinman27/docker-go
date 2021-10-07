# Inline Menu

![](../../.gitbook/assets/inline-menu.png)

{% tabs %}
{% tab title="Usage" %}
### Usage

This component allows an MSA to move throughout the Axle platform.
{% endtab %}

{% tab title="Parent Object" %}
### Parent Object

Ant Design - [Header Slider 2](https://ant.design/components/layout/)
{% endtab %}

{% tab title="Requirements" %}
### Requirements

### Behavior

* When a menu item is clicked, the appropriate item type's table should appear on the right.
* The item type that is currently open should be highlighted.
* Parent/Child nested views should be supported.

### Menu Items

* Assets - displays the list of Assets in the system in the main view.
  * This initial version will not have full system functionality.
{% endtab %}

{% tab title="Interactions" %}
### Menu Items

* If a menu item's icon or text is clicked, the right-hand view should change to show a table of the selected item type.
* If the expand arrow is clicked for a parent item, the child items should appear. The right-hand view should not be affected.
* If an expanded parent item is contracted \(by clicking the arrow\), the child items should disappear. The right-hand view should not be affected.
{% endtab %}
{% endtabs %}

