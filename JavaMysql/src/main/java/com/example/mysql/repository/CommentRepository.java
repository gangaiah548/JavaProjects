package com.example.mysql.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import com.example.mysql.model.Comment;

public interface CommentRepository extends JpaRepository<Comment, Long> {

}
