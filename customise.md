Site Go Customisation
=====================

By default you can distribute and run only the [executable](releases), but the application supports out the box customisations for:

* **Logo and favicon.ico** - need to be in */static* folder (created manually). Logo image is called 'logo.png' (can't change configuration, you'd need to change in the source-code).
Also, you need to configure *SITE_LOGO = 1* environment variable or use *-i* parameter to show the logo, which by default is not displayed.
* **Overwrite *home* and *about* pages** - you can easily overwrite *home* and *about* pages by creating new HTML files with this name (and .html extension) in the root of the website (not in */static* folder). For instance, creating *about.html* will show this file instead of the default one.
Important: HTML file needs only the main content, without header and footer, because is prepended and appended automatically with default header and footer.
* **Images** - all custom images or whatever other files (like PDF or even binary downloadable files), linked on your custom pages
need to be in */static* folder.
* **Other pages** - you can also can create new HTML pages linked from custom *home.html* or *about.html*, but they need to have all necessary content, including header and footer. They won't be available from main menu, but an inside link.
