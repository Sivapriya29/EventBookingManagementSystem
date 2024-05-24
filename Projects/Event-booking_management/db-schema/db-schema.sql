users
------
id Primary key
first_name
last_name
email unique (email, role)
password ---- Bcrypt
mobile unique (email, role)
role --- Admin/User
created_at

events
------
id Primary key
event_name
event_description
event_date
event_time
event_type
location
speaker_name
organizer_name
capacity
per_person_price
created_at
updated_at
deleted_at

bookings
--------
id
event_id
user_id
number_of_tickets
total_amount
created_at

feedbacks
----------
id
user_id
event_id
rating
comments

payments
---------
id
user_id
event_id
amount
card_number
expiry_month
expiry_year
cvv
card_holder
payment_date