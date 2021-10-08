# Calculations

## Database

| Name | Database |
| :--- | :--- |
| State / County Taxes | state\_taxes |
| City / Town Taxes | city\_taxes |
| State Title Lien Fee | state\_title\_fee |
| Plate / Tag Fee | plate\_tag\_fee |
| DMV Handling Fee | ddi\_handling\_fee |
| Other Fees | other\_fees |

## Calculations

<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Calculation</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left">Carputty Service Fee</td>
      <td style="text-align:left">(purchase_price + plate_tag_fee + state_title_fee + ddi_handling_fee +
        state_taxes + city_taxes - down_payment + other_fees + sales_tax - net_trade)*1%</td>
    </tr>
    <tr>
      <td style="text-align:left">Estimated Total Advance</td>
      <td style="text-align:left">estimated_lease_buyout_amount + plate_tag_fee + state_title_fee + ddi_handling_fee
        + state_taxes + city_taxes - down_payment + other_fees + carputty_service_fee
        + ddi_fee_elt</td>
    </tr>
    <tr>
      <td style="text-align:left">Maximum Vehicle Amount</td>
      <td style="text-align:left">
        <p>Used Vehicle Value x 110% or</p>
        <p>New Vehicle Value</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Loan-to-Value</td>
      <td style="text-align:left">Amount Total Loan / Vehicle Value</td>
    </tr>
    <tr>
      <td style="text-align:left">Taxes</td>
      <td style="text-align:left">State Taxes + City Taxes + Sales Tax</td>
    </tr>
    <tr>
      <td style="text-align:left">Tag Fee</td>
      <td style="text-align:left">Plate Tag Fee</td>
    </tr>
    <tr>
      <td style="text-align:left">Fee&apos;s</td>
      <td style="text-align:left">DMV Handling Fee + DDI Processing Fee + State Title Fee</td>
    </tr>
    <tr>
      <td style="text-align:left">Subtotal (TTF)</td>
      <td style="text-align:left">Taxes + Tag Fee + Fee&apos;s</td>
    </tr>
    <tr>
      <td style="text-align:left">Subtotal (back end)</td>
      <td style="text-align:left">Gap Insurance + Extended Contract + Roadside</td>
    </tr>
    <tr>
      <td style="text-align:left">Total Advance Amount</td>
      <td style="text-align:left">Requested Amount + Subtotal Taxes + Subtotal of Other</td>
    </tr>
    <tr>
      <td style="text-align:left">Total Loan (amount financed)</td>
      <td style="text-align:left">Asset Acquisition Fee 1% + Total Advance Amount</td>
    </tr>
  </tbody>
</table>

<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left">
        <p></p>
        <p>Line Amount</p>
      </td>
      <td style="text-align:left">
        <p></p>
        <p>Is the amount a member is approved for a transaction that ranges from
          $30,000 to $250,000.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">CP Value (Carputty Value)</td>
      <td style="text-align:left">
        <p></p>
        <p>The value Carputty assigns a vehicle that takes into account the year
          and mileage.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">
        <p></p>
        <p>Advance Amount</p>
      </td>
      <td style="text-align:left">The amount Carputty will advance towards a vehicle. For a Refi is known
        as the &quot;payoff&quot;. Retail and Lease-buyout is equal to the vehicle
        price.</td>
    </tr>
  </tbody>
</table>

### Rules

* Total Advance/loan amount cannot exceed the maximum line amount
* Initial payoff or request cannot exceed maximum vehicle amount



