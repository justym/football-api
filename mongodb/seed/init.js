db.createUser(
    {
        user: "footballer",
        pwd: "footballer",
        roles: [
            {
                role: "userAdmin",
                db: "footballDB"
            }
        ]
    }
);