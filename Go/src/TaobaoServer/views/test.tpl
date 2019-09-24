<!DOCTYPE html>
	
	<html>
	  	<head>
	    	<title>test</title>
	    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	
			<style type="text/css">
				body {
					margin: 0px;
					font-family: "Helvetica Neue",Helvetica,Arial,sans-serif;
					font-size: 14px;
					line-height: 20px;
					color: rgb(51, 51, 51);
					background-color: rgb(255, 255, 255);
				}
	
				.hero-unit {
					padding: 60px;
					margin-bottom: 30px;
					border-radius: 6px 6px 6px 6px;
				}
	
				.container {
					width: 940px;
					margin-right: auto;
					margin-left: auto;
				}
	
				.row {
					margin-left: -20px;
				}
	
				h1 {
					margin: 10px 0px;
					font-family: inherit;
					font-weight: bold;
					text-rendering: optimizelegibility;
				}
	
				.hero-unit h1 {
					margin-bottom: 0px;
					font-size: 60px;
					line-height: 1;
					letter-spacing: -1px;
					color: inherit;
				}
	
				.description {
					padding-top: 5px;
					padding-left: 5px;
					font-size: 18px;
					font-weight: 200;
					line-height: 30px;
					color: inherit;
				}
	
				p {
					margin: 0px 0px 10px;
				}
			</style>
		</head>
	
	  	<body>
	  		<header class="hero-unit" style="background-color:#A9F16C">
				<div class="container">
				<div class="row">
				  <div class="hero-text" style="background-color: whitesmoke;padding: 1em;">
				    <h1>WelCome, BlackCarDriver!</h1>
				    <p class="description">
					IP: {{.Ip}}
				    <br />Domain: {{.Domain}}
				    <br />Scheme: {{.Scheme}}
				    <br />Host: {{.Host}}
				    <br />Protocaol: {{.Protocaol}}
				    <br />Site: {{.Site}}
				    <br />UserAgent: {{.UserAgent}}
				    <br />SubDomains: {{.SubDomains}}
				    <br />Refer: {{.Refer}}
				    <br />Referer: {{.Referer}}
					<br />RunHour:{{.Runhour}}
				    </p>
				  </div>
				  <br>
				  <div style="background-color: #f0ede9;padding: 1em;">
				  <h1>Router Log</h1>
				  		<pre> {{.RouterLog}} </pre>
				  </div>
				  <br>
				  <div style="background-color: #f0ede9;padding: 1em;">
				  <h1>Models Log</h1>
				  		<pre> {{.ModelsLog}} </pre>
				  </div>
				</div>
				</div>
			</header>
		</body>
	</html>