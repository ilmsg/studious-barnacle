# studious-barnacle

    nginx 
        -> auth-api
            -> auth-grpc


/api/auth/register     post
/api/auth/login        post

/users/profile      get, post
/users/setting      get, post
/users/shops        get, post

/shops/:shopId              get
/shops/products             get
/shops/products/:productId  get