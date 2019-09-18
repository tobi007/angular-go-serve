Steps to serve your angular app from go project

1 - Create a new go mod project
2 - From the project roo run the CLI command "ng new ui"
3 - ALso to confirm that ui is correctly installed run
    "cd ui" and "ng serve --open" from the project root to start the ui.
4 - Change the outputPath in the angular.json file to "../dist"