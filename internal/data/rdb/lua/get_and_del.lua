local current = redis.call('get', KEYS[1]);
if (current) then
    redis.call('del', KEYS[1]);
end
return current