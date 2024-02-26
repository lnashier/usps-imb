# USPS Intelligent Mail (USPS 4-State Customer Code)

[![Go Report Card](https://goreportcard.com/badge/github.com/lnashier/usps-imb)](https://goreportcard.com/badge/github.com/lnashier/usps-imb)

This module provides a function named `func IMb(track, route string) (string, int)` for generating USPS Intelligent Mail barcode (IMb), also known as 4-State Customer Barcode. The Intelligent Mail barcode is used by the United States Postal Service (USPS) for sorting and tracking letters and flats, offering enhanced visibility into the mailstream.

### Parameters

- `track`: Tracking information for the mailpiece.
- `route`: Routing information for the mailpiece.

### Returns

- The generated Intelligent Mail barcode string. (len: 64)
- Status code. (0 for success)

## Structure of Intelligent Mail Barcode (IMb)

The Intelligent Mail barcode consists of the following components:

1. **Barcode ID**: A 2-digit field reserved to encode the presort identification.
    - Should be "00" if an Optional Endorsement Line (OEL) is not printed on the mailpiece.

2. **Service Type ID (STID)**: A 3-digit field used to identify the class of mail and additional services.
    - Defines the mailpiece as Full-Service, Basic, or Non-Automation.
    - Determines the disposition of Undeliverable-As-Addressed (UAA) mail.

3. **Mailer ID (MID)**: A USPS-assigned 6 or 9-digit field used to identify the mail owner or preparer.

4. **Serial Number**: A 6 or 9-digit field defined by the mailer.

5. **Routing Code**: A 5, 9, or 11-digit field identifying the delivery ZIP Code data in the address.

| Data Type     | Data Field              | # of Digits               |
|:--------------|:------------------------|:--------------------------|
| Tracking Code | Barcode Identifier      | 2 (2nd digit must be 0-4) |
|               | Service Type Identifier | 3                         |
|               | Mailer Identifier       | 6 or 9                    |
|               | Serial Number           | 6 (with 6-digit MID)      |
|               |                         | 9 (with 9-digit MID)      |
| Routing Code  | Delivery Point ZIP Code | 0, 5, 9, or 11            |
| Total         |                         | 31 maximum                |

## Mailer ID (MID) and Customer Registration ID (CRID)

- **MID (Mailer ID)**: A unique identifier assigned by USPS to Mail Owners, Mailing Agents, or service providers based on annual mail volume criteria.
    - 9-digit or 6-digit numeric code.

- **CRID (Customer Registration ID)**: A unique number identifying a specific business location involved in a mailing.
    - Required for electronic documentation submission (eDoc) and Full-Service benefits.

## STID (Service Type Identifiers)

STIDs define the mailpiece as Full-Service or Basic (non-automation) and determine the handling of UAA mail.

## Additional Resources

- [USPS 4-State Customer Code](https://barcodeguide.seagullscientific.com/Content/Symbologies/USPS_4-State_Customer_Code.htm)
- [USPS PostalPro - Intelligent Mail Barcode](https://postalpro.usps.com/onecodesolution)
- [USPS PostalPro - Intelligent Mail Barcode FAQ](https://postalpro.usps.com/node/217)
- [USPS PostalPro - Encoder/Decoder](https://postalpro.usps.com/ppro-tools/encoder-decoder)
- [USPS PostalPro - Mailer ID](https://postalpro.usps.com/mailing/mailer-id)
- [USPS PostalPro - Service Type Identifiers (STID)](https://postalpro.usps.com/service-type-identifiers/stidtable)

## License

- MIT License
- USPS License Agreement
![USPS License Agreement](usps-license.png)
