"""create publisher table

Revision ID: c49e3cfa0654
Revises: 5b2fd0d9e3b5
Create Date: 2019-06-10 18:17:42.491985

"""
import sqlalchemy as sa
from alembic import op
from sqlalchemy.sql import func

# revision identifiers, used by Alembic.
revision = 'c49e3cfa0654'
down_revision = '5b2fd0d9e3b5'
branch_labels = None
depends_on = None


def upgrade():
    op.create_table(
        'publisher',
        sa.Column('id', sa.Integer, primary_key=True),
        sa.Column('first_name', sa.String(255), nullable=False),
        sa.Column('last_name', sa.String(255), nullable=False),
        sa.Column('created_at',
                  sa.DateTime(timezone=True),
                  nullable=False,
                  server_default=func.now()),
        sa.Column('updated_at',
                  sa.DateTime(timezone=True),
                  nullable=False,
                  server_default=func.now()),
    )


def downgrade():
    op.drop_table('publisher')
