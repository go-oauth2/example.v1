<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>客户端应用OAuth2验证</title>
    <link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css">
    <link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap-theme.min.css">
    <script src="//cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
    <script src="//cdn.bootcss.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
  </head>
  <body>
    <div class ="container">
      <div class="jumbotron">
        <h1>{{.ClientName}}</h1>
        {{if .IsAuthorize}}
        <div class="alert alert-success" role="alert">
          您当前的账号已认证成功！
        </div>
        <div>
          <h2>响应令牌数据</h2>
          <pre>{{.Token}}</pre>
        </div>
        {{else}}
        <div class="alert alert-warning" role="alert">
          您当前的账号还未进行OAuth2认证，请点击认证按钮进行认证！
        </div>
        <div>
          <a role="button" class="btn btn-info btn-lg" href="{{.Href}}">认证</a>
        </div>
        {{end}}
      </div>
    </body>
  </html>