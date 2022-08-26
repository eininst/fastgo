local stockId = KEYS[1];
local decrNum = ARGV[1];
local result;
local crtStock = redis.call('get', stockId);
if crtStock == false or crtStock < decrNum then
    result = -1
else
    result = redis.call('decrBy', stockId, decrNum)
end
return result;