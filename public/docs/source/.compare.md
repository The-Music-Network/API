---
title: API Reference

language_tabs:
- bash
- javascript

includes:

search: true

toc_footers:
- <a href='http://github.com/mpociot/documentarian'>Documentation Powered by Documentarian</a>
---
<!-- START_INFO -->
# Info

Welcome to the generated API reference.
[Get Postman Collection](http://localhost/docs/collection.json)

<!-- END_INFO -->

#general
<!-- START_ec07b5f6c7252aed6c7dc3b3f3de4199 -->
## Display a listing of the resource.

> Example request:

```bash
curl -X GET "http://localhost/api/tracks" \
-H "Accept: application/json"
```

```javascript
var settings = {
    "async": true,
    "crossDomain": true,
    "url": "http://localhost/api/tracks",
    "method": "GET",
    "headers": {
        "accept": "application/json"
    }
}

$.ajax(settings).done(function (response) {
    console.log(response);
});
```

> Example response:

```json
{
    "data": [
        {
            "id": 22,
            "user_id": 9,
            "title": "Sed amet provident dolores earum voluptates illum. Quaerat iure similique optio alias occaecati accusamus quo. Aliquid quo voluptatem nisi aut laudantium incidunt consequuntur.",
            "track_url": "\/test3\/file3.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 23,
            "user_id": 6,
            "title": "Enim nam aut et sunt dolore voluptas cumque. Ut quae error enim quibusdam magnam qui. Numquam magnam sit unde error pariatur. Itaque repellendus facere explicabo inventore asperiores.",
            "track_url": "\/test4\/file4.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 24,
            "user_id": 2,
            "title": "Ullam sit at repellat qui sit similique. Dolore voluptatibus ex sapiente assumenda nobis quia.",
            "track_url": "\/test3\/file3.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 25,
            "user_id": 5,
            "title": "Et ut sunt doloremque facilis tempore illo. Est vero adipisci qui ullam asperiores dolorem. Fuga soluta provident voluptatem sed facilis minima.",
            "track_url": "\/test2\/file2.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 26,
            "user_id": 2,
            "title": "Quas at incidunt omnis quos in non pariatur sed. Voluptates saepe et quidem tenetur. Voluptatem sunt suscipit et magnam maxime et.",
            "track_url": "\/test2\/file2.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 27,
            "user_id": 5,
            "title": "Magnam ut et dolor maiores provident possimus est ex. Doloremque dolores atque iusto dicta. Ea rerum neque nulla eligendi.",
            "track_url": "\/test4\/file4.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 28,
            "user_id": 9,
            "title": "Fugit dicta et ex iste blanditiis. Sit vel exercitationem itaque voluptas et repellendus. Et dolor porro qui repellat et qui illo. Provident ut nemo fugit explicabo repellendus est.",
            "track_url": "\/test3\/file3.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 29,
            "user_id": 5,
            "title": "Aut hic aliquid voluptate non debitis. Pariatur qui doloribus est commodi officia tempora pariatur et. Voluptatem dolore explicabo nesciunt aut animi.",
            "track_url": "\/test3\/file3.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 30,
            "user_id": 4,
            "title": "Quibusdam at quo ut cum. Ut eum qui distinctio et. Dignissimos fuga ex ipsa velit aut et.",
            "track_url": "\/test2\/file2.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 31,
            "user_id": 2,
            "title": "Quia quaerat recusandae non. Dolores esse velit rerum qui et blanditiis eveniet fugiat.",
            "track_url": "\/test2\/file2.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 32,
            "user_id": 4,
            "title": "Quia maxime quia id inventore fuga. Et est aliquid et veritatis totam. Cum eaque officia ea nobis.",
            "track_url": "\/test\/file.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 33,
            "user_id": 7,
            "title": "Consequatur labore corporis eaque explicabo deserunt magnam ut illum. Voluptates iste ipsa voluptas incidunt fugit sint. Commodi quo quia nemo sunt molestias eos.",
            "track_url": "\/test4\/file4.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 34,
            "user_id": 9,
            "title": "Et vel facere doloremque dolor iure pariatur quia tempore. Quae eligendi consequuntur illum ut ut molestiae minus. Quia iste odit totam ipsa perferendis.",
            "track_url": "\/test2\/file2.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 35,
            "user_id": 9,
            "title": "Consequatur omnis sunt adipisci. Eum pariatur expedita natus aut quod. Eum eveniet maxime dolorem quo quia sit aliquid rerum. Officia ex quo earum et.",
            "track_url": "\/test3\/file3.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 36,
            "user_id": 10,
            "title": "Provident qui est dicta explicabo modi et praesentium. Numquam a harum consectetur recusandae voluptates. Corrupti exercitationem totam est sit. Quos autem eligendi quam explicabo voluptatem.",
            "track_url": "\/test1\/file1.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 37,
            "user_id": 8,
            "title": "Atque omnis laudantium quibusdam occaecati quidem id ipsum in. Accusantium autem numquam consectetur sapiente. Et delectus impedit autem animi ut.",
            "track_url": "\/test2\/file2.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 38,
            "user_id": 4,
            "title": "Et consequatur fuga possimus impedit recusandae nemo. Voluptates minus et doloribus quaerat nesciunt nostrum. Est nemo et aut consequatur voluptate.",
            "track_url": "\/test4\/file4.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 39,
            "user_id": 1,
            "title": "Quas enim rerum quidem et fugiat eos. Exercitationem voluptas aspernatur minus nihil. Nostrum nihil optio maxime sapiente similique dolores. Impedit ipsam inventore et voluptas.",
            "track_url": "\/test1\/file1.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 40,
            "user_id": 7,
            "title": "Iure dolor quisquam atque ipsum voluptas ducimus. Tenetur impedit eos itaque alias eum accusantium. Maxime nihil corrupti facilis assumenda fuga placeat qui.",
            "track_url": "\/test2\/file2.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        },
        {
            "id": 41,
            "user_id": 8,
            "title": "Incidunt sit deserunt accusamus et et sint. Labore molestiae consequuntur ad totam velit. Tempore deserunt vel ut veritatis. Ipsum soluta qui sit molestias.",
            "track_url": "\/test4\/file4.mp3",
            "created_at": "2018-08-25 03:03:41",
            "updated_at": "2018-08-25 03:03:41"
        }
    ]
}
```

### HTTP Request
`GET api/tracks`

`HEAD api/tracks`


<!-- END_ec07b5f6c7252aed6c7dc3b3f3de4199 -->

<!-- START_19885c227c1fb0a90ab68035c41da758 -->
## Show the form for creating a new resource.

> Example request:

```bash
curl -X GET "http://localhost/api/tracks/create" \
-H "Accept: application/json"
```

```javascript
var settings = {
    "async": true,
    "crossDomain": true,
    "url": "http://localhost/api/tracks/create",
    "method": "GET",
    "headers": {
        "accept": "application/json"
    }
}

$.ajax(settings).done(function (response) {
    console.log(response);
});
```

> Example response:

```json
null
```

### HTTP Request
`GET api/tracks/create`

`HEAD api/tracks/create`


<!-- END_19885c227c1fb0a90ab68035c41da758 -->

<!-- START_34e227abd1ba3b0e3bd7a7b653f81a2f -->
## Store a newly created resource in storage.

> Example request:

```bash
curl -X POST "http://localhost/api/tracks" \
-H "Accept: application/json"
```

```javascript
var settings = {
    "async": true,
    "crossDomain": true,
    "url": "http://localhost/api/tracks",
    "method": "POST",
    "headers": {
        "accept": "application/json"
    }
}

$.ajax(settings).done(function (response) {
    console.log(response);
});
```


### HTTP Request
`POST api/tracks`


<!-- END_34e227abd1ba3b0e3bd7a7b653f81a2f -->

<!-- START_e1be0325a0e533eb0a7a95083ab037f8 -->
## Display the specified resource.

> Example request:

```bash
curl -X GET "http://localhost/api/tracks/{track}" \
-H "Accept: application/json"
```

```javascript
var settings = {
    "async": true,
    "crossDomain": true,
    "url": "http://localhost/api/tracks/{track}",
    "method": "GET",
    "headers": {
        "accept": "application/json"
    }
}

$.ajax(settings).done(function (response) {
    console.log(response);
});
```

> Example response:

```json
{
    "message": "No query results for model [App\\Track].",
    "exception": "Symfony\\Component\\HttpKernel\\Exception\\NotFoundHttpException",
    "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Foundation\/Exceptions\/Handler.php",
    "line": 199,
    "trace": [
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Foundation\/Exceptions\/Handler.php",
            "line": 175,
            "function": "prepareException",
            "class": "Illuminate\\Foundation\\Exceptions\\Handler",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/app\/Exceptions\/Handler.php",
            "line": 49,
            "function": "render",
            "class": "Illuminate\\Foundation\\Exceptions\\Handler",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/nunomaduro\/collision\/src\/Adapters\/Laravel\/ExceptionHandler.php",
            "line": 68,
            "function": "render",
            "class": "App\\Exceptions\\Handler",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Routing\/Pipeline.php",
            "line": 83,
            "function": "render",
            "class": "NunoMaduro\\Collision\\Adapters\\Laravel\\ExceptionHandler",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Routing\/Pipeline.php",
            "line": 55,
            "function": "handleException",
            "class": "Illuminate\\Routing\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Routing\/Middleware\/ThrottleRequests.php",
            "line": 57,
            "function": "Illuminate\\Routing\\{closure}",
            "class": "Illuminate\\Routing\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Pipeline\/Pipeline.php",
            "line": 151,
            "function": "handle",
            "class": "Illuminate\\Routing\\Middleware\\ThrottleRequests",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Routing\/Pipeline.php",
            "line": 53,
            "function": "Illuminate\\Pipeline\\{closure}",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Pipeline\/Pipeline.php",
            "line": 104,
            "function": "Illuminate\\Routing\\{closure}",
            "class": "Illuminate\\Routing\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Routing\/Router.php",
            "line": 667,
            "function": "then",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Routing\/Router.php",
            "line": 642,
            "function": "runRouteWithinStack",
            "class": "Illuminate\\Routing\\Router",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Routing\/Router.php",
            "line": 608,
            "function": "runRoute",
            "class": "Illuminate\\Routing\\Router",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Routing\/Router.php",
            "line": 597,
            "function": "dispatchToRoute",
            "class": "Illuminate\\Routing\\Router",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Foundation\/Http\/Kernel.php",
            "line": 176,
            "function": "dispatch",
            "class": "Illuminate\\Routing\\Router",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Routing\/Pipeline.php",
            "line": 30,
            "function": "Illuminate\\Foundation\\Http\\{closure}",
            "class": "Illuminate\\Foundation\\Http\\Kernel",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/fideloper\/proxy\/src\/TrustProxies.php",
            "line": 57,
            "function": "Illuminate\\Routing\\{closure}",
            "class": "Illuminate\\Routing\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Pipeline\/Pipeline.php",
            "line": 151,
            "function": "handle",
            "class": "Fideloper\\Proxy\\TrustProxies",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Routing\/Pipeline.php",
            "line": 53,
            "function": "Illuminate\\Pipeline\\{closure}",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Foundation\/Http\/Middleware\/TransformsRequest.php",
            "line": 31,
            "function": "Illuminate\\Routing\\{closure}",
            "class": "Illuminate\\Routing\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Pipeline\/Pipeline.php",
            "line": 151,
            "function": "handle",
            "class": "Illuminate\\Foundation\\Http\\Middleware\\TransformsRequest",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Routing\/Pipeline.php",
            "line": 53,
            "function": "Illuminate\\Pipeline\\{closure}",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Foundation\/Http\/Middleware\/TransformsRequest.php",
            "line": 31,
            "function": "Illuminate\\Routing\\{closure}",
            "class": "Illuminate\\Routing\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Pipeline\/Pipeline.php",
            "line": 151,
            "function": "handle",
            "class": "Illuminate\\Foundation\\Http\\Middleware\\TransformsRequest",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Routing\/Pipeline.php",
            "line": 53,
            "function": "Illuminate\\Pipeline\\{closure}",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Foundation\/Http\/Middleware\/ValidatePostSize.php",
            "line": 27,
            "function": "Illuminate\\Routing\\{closure}",
            "class": "Illuminate\\Routing\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Pipeline\/Pipeline.php",
            "line": 151,
            "function": "handle",
            "class": "Illuminate\\Foundation\\Http\\Middleware\\ValidatePostSize",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Routing\/Pipeline.php",
            "line": 53,
            "function": "Illuminate\\Pipeline\\{closure}",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Foundation\/Http\/Middleware\/CheckForMaintenanceMode.php",
            "line": 62,
            "function": "Illuminate\\Routing\\{closure}",
            "class": "Illuminate\\Routing\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Pipeline\/Pipeline.php",
            "line": 151,
            "function": "handle",
            "class": "Illuminate\\Foundation\\Http\\Middleware\\CheckForMaintenanceMode",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Routing\/Pipeline.php",
            "line": 53,
            "function": "Illuminate\\Pipeline\\{closure}",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Pipeline\/Pipeline.php",
            "line": 104,
            "function": "Illuminate\\Routing\\{closure}",
            "class": "Illuminate\\Routing\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Foundation\/Http\/Kernel.php",
            "line": 151,
            "function": "then",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Foundation\/Http\/Kernel.php",
            "line": 116,
            "function": "sendRequestThroughRouter",
            "class": "Illuminate\\Foundation\\Http\\Kernel",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/mpociot\/laravel-apidoc-generator\/src\/Mpociot\/ApiDoc\/Generators\/LaravelGenerator.php",
            "line": 139,
            "function": "handle",
            "class": "Illuminate\\Foundation\\Http\\Kernel",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/mpociot\/laravel-apidoc-generator\/src\/Mpociot\/ApiDoc\/Generators\/AbstractGenerator.php",
            "line": 125,
            "function": "callRoute",
            "class": "Mpociot\\ApiDoc\\Generators\\LaravelGenerator",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/mpociot\/laravel-apidoc-generator\/src\/Mpociot\/ApiDoc\/Generators\/LaravelGenerator.php",
            "line": 79,
            "function": "getRouteResponse",
            "class": "Mpociot\\ApiDoc\\Generators\\AbstractGenerator",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/mpociot\/laravel-apidoc-generator\/src\/Mpociot\/ApiDoc\/Commands\/GenerateDocumentation.php",
            "line": 264,
            "function": "processRoute",
            "class": "Mpociot\\ApiDoc\\Generators\\LaravelGenerator",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/mpociot\/laravel-apidoc-generator\/src\/Mpociot\/ApiDoc\/Commands\/GenerateDocumentation.php",
            "line": 85,
            "function": "processLaravelRoutes",
            "class": "Mpociot\\ApiDoc\\Commands\\GenerateDocumentation",
            "type": "->"
        },
        {
            "function": "handle",
            "class": "Mpociot\\ApiDoc\\Commands\\GenerateDocumentation",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Container\/BoundMethod.php",
            "line": 29,
            "function": "call_user_func_array"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Container\/BoundMethod.php",
            "line": 87,
            "function": "Illuminate\\Container\\{closure}",
            "class": "Illuminate\\Container\\BoundMethod",
            "type": "::"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Container\/BoundMethod.php",
            "line": 31,
            "function": "callBoundMethod",
            "class": "Illuminate\\Container\\BoundMethod",
            "type": "::"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Container\/Container.php",
            "line": 564,
            "function": "call",
            "class": "Illuminate\\Container\\BoundMethod",
            "type": "::"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Console\/Command.php",
            "line": 179,
            "function": "call",
            "class": "Illuminate\\Container\\Container",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/symfony\/console\/Command\/Command.php",
            "line": 251,
            "function": "execute",
            "class": "Illuminate\\Console\\Command",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Console\/Command.php",
            "line": 166,
            "function": "run",
            "class": "Symfony\\Component\\Console\\Command\\Command",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/symfony\/console\/Application.php",
            "line": 886,
            "function": "run",
            "class": "Illuminate\\Console\\Command",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/symfony\/console\/Application.php",
            "line": 262,
            "function": "doRunCommand",
            "class": "Symfony\\Component\\Console\\Application",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/symfony\/console\/Application.php",
            "line": 145,
            "function": "doRun",
            "class": "Symfony\\Component\\Console\\Application",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Console\/Application.php",
            "line": 89,
            "function": "run",
            "class": "Symfony\\Component\\Console\\Application",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/vendor\/laravel\/framework\/src\/Illuminate\/Foundation\/Console\/Kernel.php",
            "line": 122,
            "function": "run",
            "class": "Illuminate\\Console\\Application",
            "type": "->"
        },
        {
            "file": "\/Users\/Jordan\/Sites\/TMN-API\/artisan",
            "line": 37,
            "function": "handle",
            "class": "Illuminate\\Foundation\\Console\\Kernel",
            "type": "->"
        }
    ]
}
```

### HTTP Request
`GET api/tracks/{track}`

`HEAD api/tracks/{track}`


<!-- END_e1be0325a0e533eb0a7a95083ab037f8 -->

<!-- START_e1fe6752c09a205ad9b5d78290e5c464 -->
## Show the form for editing the specified resource.

> Example request:

```bash
curl -X GET "http://localhost/api/tracks/{track}/edit" \
-H "Accept: application/json"
```

```javascript
var settings = {
    "async": true,
    "crossDomain": true,
    "url": "http://localhost/api/tracks/{track}/edit",
    "method": "GET",
    "headers": {
        "accept": "application/json"
    }
}

$.ajax(settings).done(function (response) {
    console.log(response);
});
```

> Example response:

```json
null
```

### HTTP Request
`GET api/tracks/{track}/edit`

`HEAD api/tracks/{track}/edit`


<!-- END_e1fe6752c09a205ad9b5d78290e5c464 -->

<!-- START_eb28c41af9bc15034fbd40d8b7651826 -->
## Update the specified resource in storage.

> Example request:

```bash
curl -X PUT "http://localhost/api/tracks/{track}" \
-H "Accept: application/json"
```

```javascript
var settings = {
    "async": true,
    "crossDomain": true,
    "url": "http://localhost/api/tracks/{track}",
    "method": "PUT",
    "headers": {
        "accept": "application/json"
    }
}

$.ajax(settings).done(function (response) {
    console.log(response);
});
```


### HTTP Request
`PUT api/tracks/{track}`

`PATCH api/tracks/{track}`


<!-- END_eb28c41af9bc15034fbd40d8b7651826 -->

<!-- START_be5d0047340eb9d8353ff10e84b21ace -->
## Remove the specified resource from storage.

> Example request:

```bash
curl -X DELETE "http://localhost/api/tracks/{track}" \
-H "Accept: application/json"
```

```javascript
var settings = {
    "async": true,
    "crossDomain": true,
    "url": "http://localhost/api/tracks/{track}",
    "method": "DELETE",
    "headers": {
        "accept": "application/json"
    }
}

$.ajax(settings).done(function (response) {
    console.log(response);
});
```


### HTTP Request
`DELETE api/tracks/{track}`


<!-- END_be5d0047340eb9d8353ff10e84b21ace -->

<!-- START_da5727be600e4865c1b632f7745c8e91 -->
## Display a listing of the resource.

> Example request:

```bash
curl -X GET "http://localhost/api/users" \
-H "Accept: application/json"
```

```javascript
var settings = {
    "async": true,
    "crossDomain": true,
    "url": "http://localhost/api/users",
    "method": "GET",
    "headers": {
        "accept": "application/json"
    }
}

$.ajax(settings).done(function (response) {
    console.log(response);
});
```

> Example response:

```json
{
    "data": [
        {
            "id": 1,
            "name": "Ramon Romaguera",
            "email": "pacocha.francesca@example.net",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 2,
            "name": "Mabelle Brakus",
            "email": "turner.alycia@example.net",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 3,
            "name": "Dr. Emiliano Kuphal",
            "email": "kub.kelvin@example.org",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 4,
            "name": "Howard Nitzsche",
            "email": "christ95@example.org",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 5,
            "name": "Everardo Boyle",
            "email": "wisoky.kory@example.com",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 6,
            "name": "Tyrese Moore",
            "email": "jgibson@example.net",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 7,
            "name": "Jeremie Moore",
            "email": "moen.laurianne@example.net",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 8,
            "name": "Magnolia Buckridge",
            "email": "ihintz@example.net",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 9,
            "name": "Dr. Carmine Wiza",
            "email": "elmer.damore@example.org",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 10,
            "name": "Sincere Buckridge Sr.",
            "email": "nyah11@example.net",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 11,
            "name": "Donato Schneider",
            "email": "elaina00@example.com",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 12,
            "name": "Dr. Llewellyn Senger",
            "email": "hartmann.kasandra@example.org",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 13,
            "name": "Breana Morar",
            "email": "hudson.jolie@example.net",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 14,
            "name": "Prof. Lesley Blanda",
            "email": "reyna71@example.com",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 15,
            "name": "Mr. Sylvan Nikolaus",
            "email": "qwintheiser@example.org",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 16,
            "name": "Annabell Metz Sr.",
            "email": "ymayert@example.net",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 17,
            "name": "Zoey Okuneva",
            "email": "cynthia24@example.com",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 18,
            "name": "Prof. Carlos Sauer",
            "email": "rowe.eugene@example.org",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 19,
            "name": "Peggie Grant",
            "email": "baylee.grant@example.com",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        },
        {
            "id": 20,
            "name": "Lexus Kris",
            "email": "jacobson.ellie@example.net",
            "created_at": "2018-08-25 03:02:00",
            "updated_at": "2018-08-25 03:02:00"
        }
    ]
}
```

### HTTP Request
`GET api/users`

`HEAD api/users`


<!-- END_da5727be600e4865c1b632f7745c8e91 -->

<!-- START_4e4753f9744661bac1222d8c1d4f2ff5 -->
## Show the form for creating a new resource.

> Example request:

```bash
curl -X GET "http://localhost/api/users/create" \
-H "Accept: application/json"
```

```javascript
var settings = {
    "async": true,
    "crossDomain": true,
    "url": "http://localhost/api/users/create",
    "method": "GET",
    "headers": {
        "accept": "application/json"
    }
}

$.ajax(settings).done(function (response) {
    console.log(response);
});
```

> Example response:

```json
null
```

### HTTP Request
`GET api/users/create`

`HEAD api/users/create`


<!-- END_4e4753f9744661bac1222d8c1d4f2ff5 -->

<!-- START_12e37982cc5398c7100e59625ebb5514 -->
## Store a newly created resource in storage.

> Example request:

```bash
curl -X POST "http://localhost/api/users" \
-H "Accept: application/json"
```

```javascript
var settings = {
    "async": true,
    "crossDomain": true,
    "url": "http://localhost/api/users",
    "method": "POST",
    "headers": {
        "accept": "application/json"
    }
}

$.ajax(settings).done(function (response) {
    console.log(response);
});
```


### HTTP Request
`POST api/users`


<!-- END_12e37982cc5398c7100e59625ebb5514 -->

<!-- START_8f99b42746e451f8dc43742e118cb47b -->
## Display the specified resource.

> Example request:

```bash
curl -X GET "http://localhost/api/users/{user}" \
-H "Accept: application/json"
```

```javascript
var settings = {
    "async": true,
    "crossDomain": true,
    "url": "http://localhost/api/users/{user}",
    "method": "GET",
    "headers": {
        "accept": "application/json"
    }
}

$.ajax(settings).done(function (response) {
    console.log(response);
});
```

> Example response:

```json
{
    "data": {
        "id": 1,
        "name": "Ramon Romaguera",
        "email": "pacocha.francesca@example.net",
        "tracks": [
            {
                "id": 39,
                "user_id": 1,
                "title": "Quas enim rerum quidem et fugiat eos. Exercitationem voluptas aspernatur minus nihil. Nostrum nihil optio maxime sapiente similique dolores. Impedit ipsam inventore et voluptas.",
                "track_url": "\/test1\/file1.mp3",
                "created_at": {
                    "date": "2018-08-25 03:03:41.000000",
                    "timezone_type": 3,
                    "timezone": "UTC"
                },
                "updated_at": {
                    "date": "2018-08-25 03:03:41.000000",
                    "timezone_type": 3,
                    "timezone": "UTC"
                }
            }
        ],
        "created_at": {
            "date": "2018-08-25 03:02:00.000000",
            "timezone_type": 3,
            "timezone": "UTC"
        },
        "updated_at": {
            "date": "2018-08-25 03:02:00.000000",
            "timezone_type": 3,
            "timezone": "UTC"
        }
    }
}
```

### HTTP Request
`GET api/users/{user}`

`HEAD api/users/{user}`


<!-- END_8f99b42746e451f8dc43742e118cb47b -->

<!-- START_6fd4489b8b8c9812aa72c8e332ce8b39 -->
## Show the form for editing the specified resource.

> Example request:

```bash
curl -X GET "http://localhost/api/users/{user}/edit" \
-H "Accept: application/json"
```

```javascript
var settings = {
    "async": true,
    "crossDomain": true,
    "url": "http://localhost/api/users/{user}/edit",
    "method": "GET",
    "headers": {
        "accept": "application/json"
    }
}

$.ajax(settings).done(function (response) {
    console.log(response);
});
```

> Example response:

```json
null
```

### HTTP Request
`GET api/users/{user}/edit`

`HEAD api/users/{user}/edit`


<!-- END_6fd4489b8b8c9812aa72c8e332ce8b39 -->

<!-- START_48a3115be98493a3c866eb0e23347262 -->
## Update the specified resource in storage.

> Example request:

```bash
curl -X PUT "http://localhost/api/users/{user}" \
-H "Accept: application/json"
```

```javascript
var settings = {
    "async": true,
    "crossDomain": true,
    "url": "http://localhost/api/users/{user}",
    "method": "PUT",
    "headers": {
        "accept": "application/json"
    }
}

$.ajax(settings).done(function (response) {
    console.log(response);
});
```


### HTTP Request
`PUT api/users/{user}`

`PATCH api/users/{user}`


<!-- END_48a3115be98493a3c866eb0e23347262 -->

<!-- START_d2db7a9fe3abd141d5adbc367a88e969 -->
## Remove the specified resource from storage.

> Example request:

```bash
curl -X DELETE "http://localhost/api/users/{user}" \
-H "Accept: application/json"
```

```javascript
var settings = {
    "async": true,
    "crossDomain": true,
    "url": "http://localhost/api/users/{user}",
    "method": "DELETE",
    "headers": {
        "accept": "application/json"
    }
}

$.ajax(settings).done(function (response) {
    console.log(response);
});
```


### HTTP Request
`DELETE api/users/{user}`


<!-- END_d2db7a9fe3abd141d5adbc367a88e969 -->

