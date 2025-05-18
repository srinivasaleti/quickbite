# Coupons Folder

This folder **must contain** the following files:

- `couponbase1.gz`
- `couponbase2.gz`
- `couponbase3.gz`

- Please add aboves file in this folder.

---

### Required for Coupon Validation

These files are **required** for the coupon filtering system to work correctly.  
Make sure they exist **before running the service**.

If **any file is missing**, the app will **not** be able to validate coupons.

---

### Important Note for Testing Coupons

The coupon files are **very large** and take time to load after the server starts.  
So when testing with coupons, **please wait 10–20 seconds** after starting the server.  
This ensures all coupons are properly loaded.

---

### Discounts on Coupon code

The coupon codes `HAPPYHOURS` and `BUYGETONE` are **not present** in the original coupon files.  
However, they are listed in the challenge description here:  
[oolio-group/kart-challenge](https://github.com/oolio-group/kart-challenge)

So I have **explicitly added support** for them:

- **`HAPPYHOURS`** – Applies an **18% discount** on the total order.
- **`BUYGETONE`** – Gives the **lowest priced item for free**.

### Note on Coupon Codes

In the challenge README here:  
[oolio-group/kart-challenge](https://github.com/oolio-group/kart-challenge/blob/advanced-challenge/backend-challenge/README.md),  
some coupon codes are mentioned under **Valid Promo Codes**, such as:

- `HAPPYHRS`
- `FIFTYOFF`

However, the README **does not provide clear details** about what these coupons should actually do.  
So, we made the following decisions:

- If the coupon code is **valid** (i.e., found in the coupon file), I apply a **10% discount** on the total order.
- If the coupon code is **invalid** (e.g., `SUPER100`), I throw an **error** saying the coupon is not valid.

This approach ensures all valid coupons have **some discount behavior**, while clearly handling invalid inputs.

If more detailed information is added in the future, the logic can be updated accordingly.
