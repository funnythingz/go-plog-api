# plog API

This is colorful pretty log aplication API.

## Usage

access with auth.

```
access_key: unkounko
```

### Plog

#### GET Plogs

```
$ curl -i /v1/plogs\?access_key=unkounko
```

#### GET Plog

```
$ curl -i /v1/plogs/:id\?access_key=unkounko
```

#### POST Plog

```
$ curl -i -F "plog[content]=hello world" -F "color[id]=2" /v1/plogs\?access_key=unkounko
```

### Color

#### GET ColorList

```
$ curl -i /v1/colors\?access_key=unkounko
```

#### POST Color (Developer only)

```
$ curl -i -F "color[name]=Blue500" -F "color[color_code]=#2196F3" -F "color[text_code]=#FFFFFF" /v1/colors\?access_key=unkounko
```

&copy; funnythingz
