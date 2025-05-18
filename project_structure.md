# High-Level Architecture
```
.github          # GitHub files like actions and workflows  
api              # OpenAPI specs and definitions  
coupons          # Folder to keep all coupons.
/server          # Backend server code  
/ui              # Frontend user interface code  
/tools           # Useful tools for the project  
  /docker        # Docker related files and configs  
  /make          # Makefiles and build scripts  
go.mod           # Go module file  
go.sum           # Go module checksum file  
Makefile         # Main makefile for build and run commands  
```

# Backend Folder Structure
```
.github              # GitHub files like actions and workflows  
api                  # OpenAPI specs and definitions  
coupons              # Keep coupons here.
server               # Main backend server code  
  /cmd               # Entry points for the server commands.  
                     # Add commands here like server or custom commands  
  /internal          # Internal code, not for outside use  
    /config          # Config files and settings  
    /database        # Database connection and helpers  
      /migrations    # Database migration files embedded in Go code  
    /domain          # Business logic organized by domain  
      /order         # Order related code  
        /handler     # HTTP handlers for order routes (handler.go)  
        /db          # Database queries for orders (db.go)  
        /service     # Order business logic layer (service.go), optional  
        routes.go    # Setup routes for order domain  
      /product       # Same structure as order, but for products  
      /coupon        # Same structure as order, but for coupons  
    /server          # Server setup and routing (server.go)  

  /pkg               # Shared reusable packages for the project  
    /bloomfilter     # Example reusable package (bloomfilter)  

  /web               # Web static files and web server code  
    /assets          # Static files like index.html  
    web.go           # Code to serve web files  

main.go              # Main entry point of the server  
```

# UI Folder Structure
```
/ui
  /node_modules       # Project dependencies  
  /public             # Public files (like index.html, icons)  
  /src                # Source code  
    /assets           # Images, fonts, and other static assets  
    /common           # Shared code like utils, hooks, API calls  
    /theme            # Theme files for styles and colors  
    /cart             # Cart feature folder  
      /components     # Cart UI components  
      /hooks          # Cart related React hooks  
    /product          # Product feature folder  
      /components     # Product UI components  
      /hooks          # Product related React hooks  
    /pages            # Pages of the app  
      /home           # Home page components and files  
    App.tsx           # Main App component  
    main.tsx          # App entry point and setup  
```