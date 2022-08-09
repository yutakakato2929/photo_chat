var query = location.search.slice(4);
if (query == "success"){
    alert("Succeed Sign Up!")
}else if (query == "passwd"){
    alert("Incorrect password.")
}else if (query == "account"){
    alert("Do not exist the account, so go to sign up page.")
}