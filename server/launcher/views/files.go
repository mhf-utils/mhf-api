package views

const FileSystemTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<title>MHF-API: Client</title>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css">
		<style>
			body {
				font-family: Arial, sans-serif;
				margin: 0;
				padding: 20px;
				background-color: #2c3e50;
			}
			h1 {
				font-size: 24px;
				color: #ecf0f1;
			}
			ul {
				list-style-type: none;
				padding: 0;
			}
			li {
				margin: 10px 0;
			}
			a {
				text-decoration: none;
				color: #ecf0f1;
				font-size: 18px;
			}
			a:hover {
				text-decoration: underline;
			}
			.icon {
				margin-right: 10px;
				font-size: 18px;
			}
			.folder {
				color: #f4a742;
			}
			.file {
				color: #90949c;
			}
			.container {
				max-width: 800px;
				margin: 0 auto;
				padding: 20px;
				background-color: #34495e;
				box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
				border-radius: 8px;
			}
			.back-link {
				margin-bottom: 20px;
				display: block;
				font-size: 18px;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>MHF-API: Client</h1>
			<ul>
				{{if not (eq .Path (printf "/%s/launcher/files" .Locale))}}
					<li>
						<a class="back-link" href="{{.ParentLink}}">
							<span class="icon folder">
								<i class="fas fa-folder"></i>
							</span>
							...
						</a>
					</li>
				{{end}}
				{{range .Files}}
				<li>
					<a href="{{.Link}}" {{if eq .Type "file"}}download{{end}}>
						<span class="icon {{if eq .Type "file"}}download{{else}}{{.Type}}{{end}}">
							{{if eq .Type "folder"}}
									<i class="fas fa-folder"></i>
							{{else}}
									<i class="fas fa-file"></i>
							{{end}}
						</span>
						{{.Name}}
					</a>
				</li>
				{{end}}
			</ul>
		</div>
	</body>
</html>

`
