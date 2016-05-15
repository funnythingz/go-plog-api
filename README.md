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
$ curl -i /api/v1/plogs\?access_key=unkounko
```

#### GET Plog

```
$ curl -i /api/v1/plogs/:id\?access_key=unkounko
```

#### POST Plog

```
$ curl -i -F "plog[content]=hello world" -F "color[color_id]=2" /api/v1/plogs\?access_key=unkounko
```

### Comment

#### GET Comments

```
$ curl -i /api/v1/Comments\?access_key=unkounko
```

#### GET Comment

```
$ curl -i /api/v1/Comments/:id\?access_key=unkounko
```

#### POST Comment

```
$ curl -i -F "comment[content]=hello world" -F "plog[plog_id]=1" /api/v1/comments\?access_key=unkounko
```

### Color

#### GET ColorList

```
$ curl -i /api/v1/colors\?access_key=unkounko
```

#### POST Color (Developer only)

```
$ curl -i -F "color[name]=Blue500" -F "color[color_code]=#2196F3" -F "color[text_code]=#FFFFFF" /api/v1/colors\?access_key=unkounko
```

&copy; funnythingz
