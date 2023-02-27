CREATE TABLE IF NOT EXISTS kbr_like(
    id BIGINT AUTO_INCREMENT,
    user_id BIGINT NULL,
    episode_id BIGINT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at BIGINT NULL DEFAULT 0,
    CONSTRAINT healty_pk PRIMARY KEY(id)
);