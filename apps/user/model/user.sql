-- 用户表
CREATE TABLE "user"
(
    id        SERIAL PRIMARY KEY,                       -- id
    name      VARCHAR(255) NULL,                         -- 用户名
    password  VARCHAR(255) NOT NULL DEFAULT '',          -- 密码
    mobile    VARCHAR(255) NOT NULL DEFAULT '',          -- 手机号码
    gender    CHAR(10) NOT NULL DEFAULT 'male',          -- 性别，默认 male
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,       -- 创建时间，默认写入插入时的时间
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,       -- 更新时间，默认写入当前时间戳
    CONSTRAINT mobile_index UNIQUE (mobile),            -- 创建唯一索引，确保 mobile 唯一
    CONSTRAINT name_index UNIQUE (name)                 -- 创建唯一索引，确保 name 唯一
);

COMMENT ON TABLE "user" IS 'user table';

-- 创建触发器确保 update_at 每次更新时自动设置当前时间
-- CREATE OR REPLACE FUNCTION update_timestamp()
-- RETURNS TRIGGER AS $$
-- BEGIN
--     NEW.update_at = CURRENT_TIMESTAMP;
-- RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;
--
-- CREATE TRIGGER update_user_update_at
--     BEFORE UPDATE ON "user"
--     FOR EACH ROW
--     EXECUTE FUNCTION update_timestamp();



