var query = location.search.slice(4);
if (query == "account"){
    alert("The account name is already exist, so you should alter a new accont name.")
}else if (query == "name"){
    alert("The User name is already exist, so you should alter a new user name.")
}else if (query == "passwd"){
    alert("Password is not matched. So you should put same passwords.")
}