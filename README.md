ticketer
A helpdesk CLI. Create tickets, list them, and close them. 

Run 

git clone https://github.com/Aloofhoneybee/ticketer.git  

cd ticketer  

go build

go install .

Example

ticketer add "fix wifi"  

ticketer add "reset password"

ticketer close 1

ticketer list

ticketer list --all
