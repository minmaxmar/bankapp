# bankapp
TODO:

0) in CreateBankClient and in CreateCard wrap in transaction to set locks on related bankID, clientID
    before creating any methods about money

1) http://127.0.0.1:3000/banks
{
        "ID": 1,
        "CreatedAt": "2025-02-18T16:57:58.863287Z",
        "UpdatedAt": "2025-02-18T16:57:58.863287Z",
        "DeletedAt": null,
        "name": "T",
        "Clients": null,
        "Cards": null
}
Clients": null 
    - while in db:
    bankapp=# SELECT * FROM bank_clients;
 client_id | bank_id
-----------+---------
         1 |       1
(1 row)

2) Add TESTS and Mock

3) specific banks with their clients should be all separate services and my app
    when trying to make a transaction should send a request to the bank identified by card number