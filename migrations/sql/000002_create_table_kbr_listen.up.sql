CREATE TABLE IF NOT EXISTS kbr_listen(
    id BIGINT AUTO_INCREMENT,
    user_id BIGINT NULL,
    session_id VARCHAR(50) NULL,
    episode_id BIGINT NULL,
    play BIGINT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at BIGINT NULL DEFAULT 0,
    CONSTRAINT kbr_listen_pk PRIMARY KEY(id)
);