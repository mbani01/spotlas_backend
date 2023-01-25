select name, substring(website from '(?:.*://)?(?:www\.)?([^/?]*)') as domain, 
COUNT(website) over (partition by website) as domain_count 
from "MY_TABLE" where "MY_TABLE".website 
in (select website from "MY_TABLE" where website <> '' group by website having count(website) > 1);