
## Cookie，Session_id and JWT

| Feature          | Cookie                         | SessionID                                | JWT                                                 |
| ---------------- | ------------------------------ | ---------------------------------------- | --------------------------------------------------- |
| ​**​Storage​**​  | Client-side (browser) <br><br> | Server-side (with ID in cookie) <br><br> | Client-side (localStorage/cookie) <br>              |
| ​**​State​**​    | Stateful (server tracks)       | Stateful (server stores data)            | Stateless (self-contained) <br>                     |
| ​**​Security​**​ | Vulnerable to CSRF/XSS <br>    | Secure (server-side data) <br>           | Secure (signed, but payload is base64-encoded) <br> |
| ​**​Use Case​**​ | Simple session tracking        | Server-rendered apps <br>                | APIs/distributed systems                            |
**Cookie/SessionID​**​: Server-dependent, simpler but less scalable.
​**​JWT​**​: Decentralized, scalable but harder to revoke

cookie is sent via Cookie header   (Vulnerable to CSRF/XSS if not secured)
Only the session ID (via cookie) is sent. Session data is stored server-side
JWT is secure ,via Cookie but not stored in server , need blacklist 



## Auth2 && SSO
SSO（Single Sign-on） ： visit TaoBao and login then you will see you also login in Tianmao   (share the session code  in storage)

OAuth2 i: kind of a protocal ,Registered in one place, it can be used in multiple places (contain "scope", "role" ...)
details:
https://www.bilibili.com/video/BV14D4y1w7dN?spm_id_from=333.788.videopod.sections&vd_source=db976053e6d6783c88dfdcd12a6212d7

